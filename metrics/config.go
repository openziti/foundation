/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package metrics

import (
	"errors"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/metrics/metrics_pb"
	"time"
)

// HandlerType is used to define known handler types
type HandlerType string

const (
	HandlerTypeInfluxDB HandlerType = "influxdb"
	HandlerTypeJSONFile HandlerType = "jsonfile"
	HandlerTypeFile     HandlerType = "file"
)

// Handler represents a sink for metric events
type Handler interface {
	// AcceptMetrics is called when new metrics become available
	AcceptMetrics(message *metrics_pb.MetricsMessage)
}

type Config struct {
	Handlers       []Handler
	Source         string
	Tags           map[string]string
	ReportInterval time.Duration
}

func LoadConfig(srcmap map[interface{}]interface{}) (*Config, error) {
	cfg := &Config{
		ReportInterval: 15 * time.Second,
	}

	pfxlog.Logger().Infof("Loading metrics configs")

	for k, v := range srcmap {
		if name, ok := k.(string); ok {
			switch name {
			case string(HandlerTypeJSONFile), string(HandlerTypeFile):
				if submap, ok := v.(map[interface{}]interface{}); ok {
					if outputFileConfig, err := LoadOutputFileConfig(submap); err == nil {
						if outputFileHandler, err := NewOutputFileMetricsHandler(outputFileConfig); err == nil {
							cfg.Handlers = append(cfg.Handlers, outputFileHandler)
							pfxlog.Logger().Infof("added metrics output file handler")
						} else {
							return nil, fmt.Errorf("error creating metrics output file handler (%s)", err)
						}
					} else {
						pfxlog.Logger().Warnf("error loading the metrics output file handler: (%s)", err)
					}
				} else {
					return nil, errors.New("invalid config for metrics output file handler ")
				}

			case string(HandlerTypeInfluxDB):
				if submap, ok := v.(map[interface{}]interface{}); ok {
					if influxCfg, err := LoadInfluxConfig(submap); err == nil {
						if influxHandler, err := NewInfluxDBMetricsHandler(influxCfg); err == nil {
							cfg.Handlers = append(cfg.Handlers, influxHandler)
							pfxlog.Logger().Infof("added influx handler")
						} else {
							return nil, fmt.Errorf("error creating influx handler (%s)", err)
						}
					}
				} else {
					return nil, errors.New("invalid influx stanza")
				}
			case "reportInterval":
				val, ok := v.(string)
				if !ok {
					return nil, errors.New("metrics.reportInterval must be a string duration, for example: 15s")
				}
				interval, err := time.ParseDuration(val)
				if err != nil {
					return nil, err
				}
				cfg.ReportInterval = interval
			case "msgQueueDepth":
				val, ok := v.(string)
				if !ok {
					return nil, errors.New("metrics.reportInterval must be a string duration, for example: 15s")
				}
				interval, err := time.ParseDuration(val)
				if err != nil {
					return nil, err
				}
				cfg.ReportInterval = interval

			}
		}
	}

	return cfg, nil
}
