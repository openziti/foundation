package metrics

import (
	"github.com/openziti/foundation/metrics/metrics_pb"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testData struct {
	registry *usageRegistryImpl
	events   []*metrics_pb.MetricsMessage
}

func setUpTest(t *testing.T) *testData {
	td := &testData{
		registry: &usageRegistryImpl{
			registryImpl: registryImpl{
				sourceId:  t.Name(),
				metricMap: cmap.New(),
			},
			intervalBucketChan: make(chan *bucketEvent, 1),
		}}
	td.registry.eventSink = td
	return td
}

func (t *testData) AddHandler(handler Handler) {
	panic("implement me")
}

func (t *testData) RemoveHandler(handler Handler) {
	panic("implement me")
}

func (t *testData) Shutdown() {
	panic("implement me")
}

func (t *testData) AcceptMetrics(e *metrics_pb.MetricsMessage) {
	t.events = append(t.events, e)
}

func TestEmpty(t *testing.T) {
	td := setUpTest(t)
	td.registry.report()
	assert.Len(t, td.events, 0)
}

func Test_Histogram(t *testing.T) {
	td := setUpTest(t)

	hist := td.registry.Histogram("test.hist")
	hist.Update(10)

	td.registry.report()
	assert.Len(t, td.events, 1)

	ev := td.events[0]
	assert.Nil(t, ev.FloatValues)
	assert.Nil(t, ev.Meters)
	assert.Nil(t, ev.IntValues)

	assert.NotNil(t, ev.Histograms)

	hm := ev.Histograms["test.hist"]
	assert.NotNil(t, hm)
	assert.Equal(t, int64(10), hm.Min)
	assert.Equal(t, int64(10), hm.Max)
}

func Test_Timer(t *testing.T) {
	td := setUpTest(t)

	timer := td.registry.Timer("test.timer")

	timer.Update(3 * time.Second)

	timer.Time(func() {
		<-time.After(time.Second)
	})

	td.registry.report()
	assert.Len(t, td.events, 1)

	ev := td.events[0]
	assert.Nil(t, ev.FloatValues)
	assert.Nil(t, ev.Meters)
	assert.Nil(t, ev.IntValues)

	hm := ev.Timers["test.timer"]
	assert.NotNil(t, hm)
	assert.Equal(t, int64(2), hm.Count)

	assert.Equal(t, 3*time.Second, time.Duration(hm.Max))
	assert.InDelta(t, time.Second, time.Duration(hm.Min), float64(time.Millisecond))
}
