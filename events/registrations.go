package events

import (
	"github.com/openziti/foundation/metrics"
	"github.com/openziti/foundation/util/cowslice"
)

func AddMetricsEventHandler(handler metrics.Handler) {
	cowslice.Append(metrics.EventHandlerRegistry, handler)
}

func RemoveMetricsEventHandler(handler metrics.Handler) {
	cowslice.Delete(metrics.EventHandlerRegistry, handler)
}
