package metrics

import (
	"fmt"
	"github.com/openziti/foundation/metrics/metrics_pb"
	cmap "github.com/orcaman/concurrent-map"
	"reflect"
	"time"
)

// UsageRegistry extends registry to allow collecting usage metrics
type UsageRegistry interface {
	Registry
	IntervalCounter(name string, intervalSize time.Duration) IntervalCounter
	SetEventSink(eventSink Handler)
	Flush()
}

func NewUsageRegistryFromConfig(config *Config, closeNotify chan struct{}) UsageRegistry {
	if config.ReportInterval == 0 {
		config.ReportInterval = 15 * time.Second
	}
	return NewUsageRegistry(config.Source, config.Tags, config.ReportInterval, config.EventSink, closeNotify)
}

func NewUsageRegistry(sourceId string, tags map[string]string, reportInterval time.Duration, eventSink Handler, closeNotify <-chan struct{}) UsageRegistry {
	registry := &usageRegistryImpl{
		registryImpl: registryImpl{
			sourceId:  sourceId,
			tags:      tags,
			metricMap: cmap.New(),
		},
		eventSink:          eventSink,
		intervalBucketChan: make(chan *bucketEvent, 1),
		closeNotify:        closeNotify,
	}

	go registry.run(reportInterval)

	return registry
}

type bucketEvent struct {
	interval *metrics_pb.MetricsMessage_IntervalCounter
	name     string
}

type usageRegistryImpl struct {
	registryImpl
	eventSink          Handler
	intervalBucketChan chan *bucketEvent
	intervalBuckets    []*bucketEvent
	closeNotify        <-chan struct{}
}

func (registry *usageRegistryImpl) SetEventSink(eventSink Handler) {
	registry.eventSink = eventSink
}

// NewIntervalCounter creates an IntervalCounter
func (registry *usageRegistryImpl) IntervalCounter(name string, intervalSize time.Duration) IntervalCounter {
	metric, present := registry.metricMap.Get(name)
	if present {
		intervalCounter, ok := metric.(IntervalCounter)
		if !ok {
			panic(fmt.Errorf("metric '%v' already exists and is not an interval counter. It is a %v", name, reflect.TypeOf(metric).Name()))
		}
		return intervalCounter
	}

	disposeF := func() { registry.dispose(name) }
	intervalCounter := newIntervalCounter(name, intervalSize, registry, time.Minute, time.Second*80, disposeF, registry.closeNotify)
	registry.metricMap.Set(name, intervalCounter)
	return intervalCounter
}

func (registry *usageRegistryImpl) Poll() *metrics_pb.MetricsMessage {
	base := registry.registryImpl.Poll()
	if base == nil && registry.intervalBuckets == nil {
		return nil
	}

	var builder *messageBuilder
	if base == nil {
		builder = newMessageBuilder(registry.sourceId, registry.tags)
	} else {
		builder = (*messageBuilder)(base)
	}

	builder.addIntervalBucketEvents(registry.intervalBuckets)
	registry.intervalBuckets = nil

	return (*metrics_pb.MetricsMessage)(builder)
}

func (registry *usageRegistryImpl) run(reportInterval time.Duration) {
	ticker := time.NewTicker(reportInterval)
	defer ticker.Stop()

	for {
		select {
		case interval := <-registry.intervalBucketChan:
			registry.intervalBuckets = append(registry.intervalBuckets, interval)
		case <-ticker.C:
			registry.report()
		case <-registry.closeNotify:
			registry.DisposeAll()
			return
		}
	}
}
func (registry *usageRegistryImpl) reportInterval(counter *intervalCounterImpl, intervalStartUTC int64, values map[string]uint64) {
	bucket := &metrics_pb.MetricsMessage_IntervalBucket{
		IntervalStartUTC: intervalStartUTC,
		Values:           values,
	}

	interval := &metrics_pb.MetricsMessage_IntervalCounter{
		IntervalLength: uint64(counter.intervalSize.Seconds()),
		Buckets:        []*metrics_pb.MetricsMessage_IntervalBucket{bucket},
	}

	bucketEvent := &bucketEvent{
		interval: interval,
		name:     counter.name,
	}

	registry.intervalBucketChan <- bucketEvent
}

func (registry *usageRegistryImpl) report() {
	if msg := registry.Poll(); msg != nil {
		registry.eventSink.AcceptMetrics(msg)
	}
}

func (registry *usageRegistryImpl) Flush() {
	registry.EachMetric(func(name string, metric Metric) {
		if ic, ok := metric.(*intervalCounterImpl); ok {
			ic.flush()
		}
	})
	time.Sleep(250 * time.Millisecond)
	registry.report()
}
