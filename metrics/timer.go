package metrics

import (
	"github.com/rcrowley/go-metrics"
	"time"
)

// Histogram represents a metric which is measuring the distribution of values for some measurement
type Timer interface {
	Metric
	// Clear()
	Update(time.Duration)
	Time(func())
}

type timerImpl struct {
	metrics.Timer
	dispose func()
}

func (t *timerImpl) Dispose() {
	t.Stop()
	t.dispose()
}
