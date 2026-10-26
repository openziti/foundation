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
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReplaceAttrRenamesLevelLowercase(t *testing.T) {
	cases := []struct {
		lvl  slog.Level
		want string
	}{
		{LevelTrace, "trace"},
		{slog.LevelDebug, "debug"},
		{slog.LevelInfo, "info"},
		{slog.LevelWarn, "warn"},
		{slog.LevelError, "error"},
		{LevelFatal, "fatal"},
		{LevelPanic, "panic"},
	}
	for _, c := range cases {
		got := ReplaceAttr(nil, slog.Any(slog.LevelKey, c.lvl))
		require.Equal(t, slog.LevelKey, got.Key)
		require.Equal(t, c.want, got.Value.String())
	}
}

func TestReplaceAttrSuppressesSourceAtTopLevel(t *testing.T) {
	src := &slog.Source{Function: "f", File: "f.go", Line: 1}
	got := ReplaceAttr(nil, slog.Any(slog.SourceKey, src))
	require.Equal(t, slog.Attr{}, got, "source should be suppressed")
}

// TestReplaceAttrIgnoresInsideGroups proves the rename/suppression only
// applies to top-level attrs.
func TestReplaceAttrIgnoresInsideGroups(t *testing.T) {
	a := slog.Any(slog.LevelKey, slog.LevelInfo)
	got := ReplaceAttr([]string{"g"}, a)
	require.Equal(t, a.Key, got.Key)
	require.Equal(t, a.Value, got.Value)
}

func TestReplaceAttrPassesThroughOtherAttrs(t *testing.T) {
	a := slog.String("k", "v")
	got := ReplaceAttr(nil, a)
	require.Equal(t, a, got)
}

// TestSourceFlattenerPassesThroughWhenPCZero proves records with no PC pass
// through the flattener without gaining file/func attrs.
func TestSourceFlattenerPassesThroughWhenPCZero(t *testing.T) {
	rec := &recordingHandler{}
	flat := &sourceFlattener{parent: rec}

	r := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
	r.AddAttrs(slog.String("k", "v"))
	require.NoError(t, flat.Handle(context.Background(), r))

	require.Equal(t, 1, rec.count())
	attrs := map[string]any{}
	rec.snapshot()[0].Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()
		return true
	})
	require.NotContains(t, attrs, "file", "no file attr should be added when PC is zero")
	require.NotContains(t, attrs, "func")
	require.Equal(t, "v", attrs["k"])
}

// TestJSONHandlerEmitsPfxlogShape proves a direct slog call produces JSON with
// a lowercase level, flat file/func keys, no nested source object, and the
// caller's attrs preserved.
func TestJSONHandlerEmitsPfxlogShape(t *testing.T) {
	buf := &bytes.Buffer{}
	slog.New(NewJSONHandler(buf)).Info("hello", "k", "v")

	line := strings.TrimSpace(buf.String())
	var got map[string]any
	require.NoError(t, json.Unmarshal([]byte(line), &got), "raw=%q", line)

	require.Equal(t, "hello", got["msg"])
	require.Equal(t, "info", got["level"], "level must be lowercase")
	require.Equal(t, "v", got["k"])
	require.NotContains(t, got, "source", "nested source must be suppressed")

	file, _ := got["file"].(string)
	require.Contains(t, file, "format_test.go:", "file should be filename:line")
	require.Equal(t, "github.com/openziti/foundation/v2/logging.TestJSONHandlerEmitsPfxlogShape", got["func"])
}

// TestJSONHandlerEmitsCustomLevels confirms LevelTrace, LevelFatal, and
// LevelPanic render with their canonical lowercase names rather than slog's
// offset form ("DEBUG-4" / "ERROR+4" / "ERROR+8").
func TestJSONHandlerEmitsCustomLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := slog.New(NewJSONHandler(buf))
	logger.Log(context.Background(), LevelTrace, "t")
	logger.Log(context.Background(), LevelFatal, "f")
	logger.Log(context.Background(), LevelPanic, "p")

	var levels []string
	for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
		var m map[string]any
		require.NoError(t, json.Unmarshal([]byte(line), &m))
		levels = append(levels, m["level"].(string))
	}
	require.Equal(t, []string{"trace", "fatal", "panic"}, levels)
}

// TestTextHandlerEmitsKeyValue proves NewTextHandler produces plain key=value
// output with the canonical lowercase level and flat caller keys.
func TestTextHandlerEmitsKeyValue(t *testing.T) {
	buf := &bytes.Buffer{}
	slog.New(NewTextHandler(buf)).Info("hello", "k", "v")

	line := strings.TrimSpace(buf.String())
	require.Contains(t, line, "level=info")
	require.Contains(t, line, "msg=hello")
	require.Contains(t, line, "k=v")
	require.Contains(t, line, "func=github.com/openziti/foundation/v2/logging.TestTextHandlerEmitsKeyValue")
	require.NotContains(t, line, "source=")
	require.False(t, json.Valid([]byte(line)), "text output must not be JSON")
}
