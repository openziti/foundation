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
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"
)

// SyncEmit writes r through the default Registry's root handler synchronously
// on the caller's goroutine. If the root is an *AsyncHandler, the call routes
// through its SyncEmit, which flushes the queued records and then writes r
// under the same mutex the drain uses, so the buffered context leading up to
// the call survives a process exit right after. Otherwise it falls back to
// Handle, which is already synchronous for any non-async handler.
//
// It is the durability primitive behind Fatal, and is available to callers that
// need the same guarantee for a record they build themselves.
func SyncEmit(ctx context.Context, r slog.Record) error {
	root := DefaultRegistry().Root()
	if ah, ok := root.(*AsyncHandler); ok {
		return ah.SyncEmit(ctx, r)
	}
	return root.Handle(ctx, r)
}

// osExit is os.Exit, indirected so tests can exercise Fatal without
// terminating the test process.
var osExit = os.Exit

// Fatal writes msg at LevelFatal and then exits the process with status 1. It
// is the slog-world equivalent of logrus.Fatal, which slog does not provide
// (slog has only Debug/Info/Warn/Error and never exits the process itself).
//
// The record is emitted durably through SyncEmit, so it flushes any queued
// records and writes synchronously before the exit, and it bypasses level
// gating so a fatal is never filtered out. Hard-exit paths call this instead
// of logging an Error and then calling os.Exit or panic: those drop the record
// because the async queue never drains before the process is gone.
func Fatal(ctx context.Context, msg string, attrs ...slog.Attr) {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip runtime.Callers + this frame
	r := slog.NewRecord(time.Now(), LevelFatal, msg, pcs[0])
	r.AddAttrs(attrs...)
	_ = SyncEmit(ctx, r)
	osExit(1)
}
