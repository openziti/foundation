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
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/metrics/metrics_pb"
	"github.com/orcaman/concurrent-map"
	"github.com/pkg/errors"
	"github.com/rcrowley/go-metrics"
	"reflect"
	"time"
)

// Metric is the base functionality for all metrics types
type Metric interface {
	Dispose()
}

// Registry allows for configuring and accessing metrics for a fabric application
type Registry interface {
	SourceId() string
	Meter(name string) Meter
	Histogram(name string) Histogram
	Timer(name string) Timer
	IntervalCounter(name string, intervalSize time.Duration) IntervalCounter
	EachMetric(visitor func(name string, metric Metric))
}

func NewRegistryFromConfig(config *Config) Registry {
	if config.ReportInterval == 0 {
		config.ReportInterval = 15 * time.Second
	}
	return NewRegistry(config.Source, config.Tags, config.ReportInterval, config.EventSink)
}

func NewRegistry(sourceId string, tags map[string]string, reportInterval time.Duration, eventSink Handler) Registry {
	registry := &registryImpl{
		sourceId:           sourceId,
		metricMap:          cmap.New(),
		eventSink:          eventSink,
		tags:               tags,
		intervalBucketChan: make(chan *bucketEvent, 1),
	}

	go registry.run(reportInterval)

	return registry
}

type bucketEvent struct {
	interval *metrics_pb.MetricsMessage_IntervalCounter
	name     string
}

type registryImpl struct {
	sourceId           string
	tags               map[string]string
	metricMap          cmap.ConcurrentMap
	eventSink          Handler
	intervalBucketChan chan *bucketEvent
	intervalBuckets    []*bucketEvent
}

func (registry *registryImpl) run(reportInterval time.Duration) {
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

func (registry *registryImpl) dispose(name string) {
	registry.metricMap.Remove(name)
}

func (registry *registryImpl) SourceId() string {
	return registry.sourceId
}

func (registry *registryImpl) Meter(name string) Meter {
	metric, present := registry.metricMap.Get(name)
	if present {
		meter, ok := metric.(Meter)
		if !ok {
			panic(fmt.Errorf("metric '%v' already exists and is not a meter. It is a %v", name, reflect.TypeOf(metric).Name()))
		}
		return meter
	}

	meter := &meterImpl{
		Meter: metrics.NewMeter(),
		dispose: func() {
			registry.dispose(name)
		},
	}
	registry.metricMap.Set(name, meter)
	return meter
}

func (registry *registryImpl) Histogram(name string) Histogram {
	metric, present := registry.metricMap.Get(name)
	if present {
		histogram, ok := metric.(Histogram)
		if !ok {
			panic(fmt.Errorf("metric '%v' already exists and is not a histogram. It is a %v", name, reflect.TypeOf(metric).Name()))
		}
		return histogram
	}

	histogram := &histogramImpl{
		Histogram: metrics.NewHistogram(metrics.NewExpDecaySample(128, 0.015)),
		dispose: func() {
			registry.dispose(name)
		},
	}
	registry.metricMap.Set(name, histogram)
	return histogram
}

func (registry *registryImpl) Timer(name string) Timer {
	metric, present := registry.metricMap.Get(name)
	if present {
		timer, ok := metric.(Timer)
		if !ok {
			panic(fmt.Errorf("metric '%v' already exists and is not a timer. It is a %v", name, reflect.TypeOf(metric).Name()))
		}
		return timer
	}

	timer := &timerImpl{
		Timer: metrics.NewTimer(),
		dispose: func() {
			registry.dispose(name)
		},
	}
	registry.metricMap.Set(name, timer)
	return timer
}

// NewIntervalCounter creates an IntervalCounter
func (registry *registryImpl) IntervalCounter(name string, intervalSize time.Duration) IntervalCounter {
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

func (registry *registryImpl) EachMetric(visitor func(name string, metric Metric)) {
	for entry := range registry.metricMap.IterBuffered() {
		visitor(entry.Key, entry.Val.(Metric))
	}
}

func (registry *registryImpl) Each(visitor func(string, interface{})) {
	for entry := range registry.metricMap.IterBuffered() {
		visitor(entry.Key, entry.Val)
	}
}

func (registry *registryImpl) reportInterval(counter *intervalCounterImpl, intervalStartUTC int64, values map[string]uint64) {
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

func (registry *registryImpl) report() {
	// If there's nothing to report, skip it
	if registry.metricMap.Count() == 0 && len(registry.intervalBuckets) == 0 {
		return
	}

	builder := newMessageBuilder(registry.sourceId, registry.tags)

	registry.EachMetric(func(name string, i Metric) {
		switch metric := i.(type) {
		case *meterImpl:
			builder.addMeter(name, metric.Snapshot())
		case *histogramImpl:
			builder.addHistogram(name, metric.Snapshot())
		case *timerImpl:
			builder.addTimer(name, metric.Snapshot())
		case *intervalCounterImpl:
			// ignore, handled below
		default:
			pfxlog.Logger().Errorf("Unsupported metric type %v", reflect.TypeOf(i))
		}
	})

	builder.addIntervalBucketEvents(registry.intervalBuckets)
	registry.intervalBuckets = nil

	msg := (*metrics_pb.MetricsMessage)(builder)
	registry.eventSink.AcceptMetrics(msg)
}

// Provide rest of go-metrics Registry interface, so we can use go-metrics reporters if desired
func (registry *registryImpl) Get(s string) interface{} {
	val, _ := registry.metricMap.Get(s)
	return val
}

func (registry *registryImpl) GetAll() map[string]map[string]interface{} {
	return nil
}

func (registry *registryImpl) GetOrRegister(s string, i interface{}) interface{} {
	return registry.metricMap.Upsert(s, i, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
		if exist {
			return valueInMap
		}
		return newValue
	})
}

func (registry *registryImpl) Register(s string, i interface{}) error {
	if registry.metricMap.SetIfAbsent(s, i) {
		return errors.Errorf("duplicate metric %v", s)
	}
	return nil
}

func (registry *registryImpl) RunHealthchecks() {
}

func (registry *registryImpl) Unregister(s string) {
	registry.metricMap.Remove(s)
}

func (registry *registryImpl) UnregisterAll() {
	for _, key := range registry.metricMap.Keys() {
		registry.Unregister(key)
	}
}
