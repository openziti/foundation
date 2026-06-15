/*
	Copyright NetFoundry Inc.

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

package logging

import (
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

// Default AsyncOptions values. See doc/design/logging-refactor.md.
const (
	DefaultQueueSize       = 4096
	DefaultBlockThreshold  = slog.LevelWarn
	DefaultSummaryInterval = 5 * time.Second
)

// AsyncOptions configures the AsyncHandler queue, block-threshold, and summary
// cadence. Zero values are not valid; use DefaultOptions as a starting point.
type AsyncOptions struct {
	// QueueSize is the bounded capacity of the records queue between the
	// handler's caller and the drain goroutine.
	QueueSize int

	// BlockThreshold is the lowest level at which Handle blocks when the queue
	// is full. Records strictly below this level are dropped under saturation
	// and counted toward the next summary line.
	BlockThreshold slog.Level

	// SummaryInterval is the cadence at which the drain emits a drop-summary
	// record when any per-level drop counter is non-zero.
	SummaryInterval time.Duration
}

// DefaultOptions returns AsyncOptions with the defaults documented in
// doc/design/logging-refactor.md.
func DefaultOptions() AsyncOptions {
	return AsyncOptions{
		QueueSize:       DefaultQueueSize,
		BlockThreshold:  DefaultBlockThreshold,
		SummaryInterval: DefaultSummaryInterval,
	}
}

// Validate returns an error if any field is outside its valid range.
func (o AsyncOptions) Validate() error {
	if o.QueueSize < 1 {
		return errors.Errorf("QueueSize must be >= 1, got %d", o.QueueSize)
	}
	if o.BlockThreshold < LevelTrace || o.BlockThreshold > LevelPanic {
		return errors.Errorf("BlockThreshold %v is outside the canonical level range (%v..%v)", o.BlockThreshold, LevelTrace, LevelPanic)
	}
	if o.SummaryInterval <= 0 {
		return errors.Errorf("SummaryInterval must be > 0, got %v", o.SummaryInterval)
	}
	return nil
}
