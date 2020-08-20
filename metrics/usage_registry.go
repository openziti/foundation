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
}

func NewUsageRegistryFromConfig(config *Config) UsageRegistry {
	if config.ReportInterval == 0 {
		config.ReportInterval = 15 * time.Second
	}
	return NewUsageRegistry(config.Source, config.Tags, config.ReportInterval, config.EventSink)
}

func NewUsageRegistry(sourceId string, tags map[string]string, reportInterval time.Duration, eventSink Handler) UsageRegistry {
	registry := &usageRegistryImpl{
		registryImpl: registryImpl{
			sourceId:  sourceId,
			tags:      tags,
			metricMap: cmap.New(),
		},
		eventSink:          eventSink,
		intervalBucketChan: make(chan *bucketEvent, 1),
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
	intervalCounter := newIntervalCounter(name, intervalSize, registry, time.Minute, time.Second*80, disposeF)
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
	ticker := time.Tick(reportInterval)

	for {
		select {
		case interval := <-registry.intervalBucketChan:
			registry.intervalBuckets = append(registry.intervalBuckets, interval)
		case <-ticker:
			registry.report()
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
