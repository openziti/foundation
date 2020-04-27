package metrics

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/netfoundry/ziti-foundation/metrics/metrics_pb"
	"errors"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"strings"
)


type jsonfileReporter struct {
	path         string
	maxsize   int
	metricsChan chan *metrics_pb.MetricsMessage
}

func (reporter *jsonfileReporter) AcceptMetrics(message *metrics_pb.MetricsMessage) {
	reporter.metricsChan <- message
}

// Message handler that write node and link metrics to a file in json format
func NewJsonFileMetricsHandler(cfg *jsonfileConfig) (Handler, error) {
	rep := &jsonfileReporter{
		path:         cfg.path,
		maxsize:    cfg.maxsizemb,
		metricsChan: make(chan *metrics_pb.MetricsMessage, 10),
	}

	go rep.run()
	return rep, nil
}


func (reporter *jsonfileReporter) run() {
	log := pfxlog.Logger()
	log.Info("JSON File Metrics Reporter started")
	defer log.Warn("exited")

	for {
		select {
		case msg := <-reporter.metricsChan:
			if err := reporter.send(msg); err != nil {
				log.Printf("unable to send metrics to JSON File handler. err=%v", err)
			}

		}
	}
}

func (reporter *jsonfileReporter) send(msg *metrics_pb.MetricsMessage) error {
	links := make(map[string]interface{})
	event := make(map[string]interface{})
	tags := make(map[string]string)

	// inject tags if there are any
	for k, v := range msg.Tags {
		tags[k] = v
	}

	event["tags"] = tags
	event["sourceId"] = msg.SourceId
	event["timestamp"] = msg.Timestamp.Seconds * 1000
	event["sourceType"] = metrics_pb.MetricsSourceType_name[int32(msg.SourceType)]

	// ints
	for name, val := range msg.IntValues {
		// skip empty objects
		if string(val) == "" {
			continue
		}

		// if link is in the name, add it to the link event
		if strings.Contains(name, "link") {
			links = ProcessLinkEvent(links, name, val)
			continue
		}

		event[name] = val
	}

	// floats
	for name, val := range msg.FloatValues {
		// if link is in the name, add it to the link event
		if strings.Contains(name, "link") {

			links = ProcessLinkEvent(links, name, val)
			continue
		}

		event[name] = val
	}

	// histograms
	for name, val := range msg.Histograms {
		// skip empty objects
		if val.String() == "" {
			continue
		}

		// if link is in the name, add it to the link event
		if strings.Contains(name, "link") {
			links = ProcessLinkEvent(links, name, val)
			continue
		}

		event[name] = val
	}

	// meters
	for name, val := range msg.Meters {
		// skip empty objects
		if val.String() == "" {
			continue
		}

		// if link is in the name, add it to the link event
		if strings.Contains(name, "link") {
			links = ProcessLinkEvent(links, name, val)
			continue
		}

		event[name] = val
	}

	// intervals
	for name, val := range msg.IntervalCounters {
		// skip empty objects
		if val.String() == "" {
			continue
		}

		// if link is in the name, add it to the link event
		if strings.Contains(name, "link") {
			links = ProcessLinkEvent(links, name, val)
			continue
		}

		event[name] = val
	}

	// json format
	out, err := json.Marshal(event)

	// Check for max file size, truncate if larger than threshold
	stat, err := os.Stat(reporter.path)
	if err != nil {
		log.Println(err)
	}
	// get the size
	size := stat.Size()
	if size >= int64(reporter.maxsize * 1024 * 1024) {
		err := os.Truncate(reporter.path, 0)
		if err != nil {
			log.Println(err)
		}
	}

	f, err := os.OpenFile(reporter.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Println(err)
	}
	// log the router-level event
	if _, err := fmt.Fprintf(f, "%s\n", out); err != nil {
		log.Println(err)
	}

	for linkName, linkEvent := range links {
		// build an event just for the link that can be rolled up by router id
		linkEvent.(map[string]interface{})["linkId"] = linkName
		linkEvent.(map[string]interface{})["sourceId"] = msg.SourceId
		linkEvent.(map[string]interface{})["timestamp"] = msg.Timestamp.Seconds * 1000
		linkEvent.(map[string]interface{})["sourceType"] = metrics_pb.MetricsSourceType_name[int32(msg.SourceType)]
		linkEvent.(map[string]interface{})["tags"] = tags

		// json format
		out, err := json.Marshal(linkEvent)
		if err != nil {
			panic (err)
		}

		// log the link-level event
		if _, err := fmt.Fprintf(f, "%s\n", out); err != nil {
			log.Println(err)
		}

	}

	defer f.Close()

	return err
}


func ProcessLinkEvent(links map[string]interface{}, name string, val interface{}) map[string]interface{}  {
	parts := strings.Split(name, ".")
	linkName := string(parts[1])

	//rename the metric key without the id in it
	metricName := strings.Replace(name, linkName+".","", 1)

	// see if the link event exists, if not create it
	_, exists := links[linkName]

	if !exists {
		links[linkName] = make(map[string]interface{})
	}

	links[linkName].(map[string]interface{})[metricName] = val

	return links
}

type jsonfileConfig struct {
	path      string
	maxsizemb   int
}


func LoadJSONFileConfig(src map[interface{}]interface{}) (*jsonfileConfig, error) {
	cfg := &jsonfileConfig{}

	if value, found := src["path"]; found {
		if path, ok := value.(string); ok {
			// check path writablility
			cfg.path = path

			f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
			if err != nil {
				return nil, fmt.Errorf("cannot write to log file path: %s", path)
			}else {
				f.Close()
			}
		} else {
			return nil, errors.New("invalid jsonFileReporter 'path' value")
		}
	} else {
		return nil, errors.New("missing required 'path' config for JSONFileReporter")
	}

	if value, found := src["maxsizemb"]; found {
		if maxsizemb, ok := value.(int); ok {
			cfg.maxsizemb = maxsizemb
		} else {
			// just set a default
			return nil, errors.New("invalid 'maxsizemb' config for JSONFileReporter")
		}
	} else {
		return nil, errors.New("missing jsonFileReporter 'maxsizemb' config")
	}

	return cfg, nil
}

