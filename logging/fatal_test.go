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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestSyncEmitDeliversThroughAsyncRoot proves the package-level SyncEmit routes
// through the *AsyncHandler root synchronously: the record reaches the
// downstream before SyncEmit returns, without waiting for the async drain.
func TestSyncEmitDeliversThroughAsyncRoot(t *testing.T) {
	resetDefaultForTest()
	rec := &recordingHandler{}
	async, err := NewAsyncHandler(rec, DefaultOptions())
	require.NoError(t, err)
	defer func() { _ = async.Close() }()
	Configure(async)

	r := slog.NewRecord(time.Now(), slog.LevelInfo, "sync-msg", 0)
	require.NoError(t, SyncEmit(context.Background(), r))

	require.Equal(t, 1, rec.count(), "SyncEmit must write before returning, not leave the record queued")
	require.Equal(t, "sync-msg", rec.snapshot()[0].Message)
}

// TestSyncEmitFallsBackForNonAsyncRoot proves SyncEmit handles a Registry whose
// root is not an *AsyncHandler by falling back to Handle. Handle on a non-async
// handler is already synchronous, so durability is preserved.
func TestSyncEmitFallsBackForNonAsyncRoot(t *testing.T) {
	resetDefaultForTest()
	rec := &recordingHandler{}
	Configure(rec) // root is the plain recording handler, not an AsyncHandler

	r := slog.NewRecord(time.Now(), LevelFatal, "fatal-fallback", 0)
	require.NoError(t, SyncEmit(context.Background(), r))
	require.Equal(t, 1, rec.count())
}

// TestFatalEmitsDurablyAndExits proves Fatal writes its record synchronously
// (present before any Close, so it survives a process exit), carries the attrs
// and caller PC, and calls osExit(1). The global level is set above Fatal to
// also prove Fatal bypasses level gating, matching logrus.Fatal.
func TestFatalEmitsDurablyAndExits(t *testing.T) {
	resetDefaultForTest()
	rec := &recordingHandler{}
	async, err := NewAsyncHandler(rec, DefaultOptions())
	require.NoError(t, err)
	defer func() { _ = async.Close() }()
	Configure(async)
	SetGlobalLevel(LevelPanic) // above Fatal: a gated path would drop it

	var gotCode int
	var exited bool
	prev := osExit
	osExit = func(code int) { gotCode = code; exited = true }
	defer func() { osExit = prev }()

	Fatal(context.Background(), "boom", slog.String("k", "v"))

	require.True(t, exited, "Fatal must call osExit")
	require.Equal(t, 1, gotCode)
	require.Equal(t, 1, rec.count(), "record must be written synchronously, not left in the queue")

	got := rec.snapshot()[0]
	require.Equal(t, LevelFatal, got.Level)
	require.Equal(t, "boom", got.Message)
	require.NotZero(t, got.PC, "Fatal should capture the caller PC")

	attrs := map[string]string{}
	got.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.String()
		return true
	})
	require.Equal(t, "v", attrs["k"])
}
