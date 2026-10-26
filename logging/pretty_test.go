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
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func plainPrettyOptions() *PrettyOptions {
	return &PrettyOptions{
		StartTimestamp:  time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		TimestampFormat: "2006-01-02 15:04:05.000",
		UseColor:        false,
	}
}

// TestPrettyHandlerLabelsAllSevenLevels proves every canonical level,
// including the custom Trace, Fatal and Panic values, renders its pfxlog label.
func TestPrettyHandlerLabelsAllSevenLevels(t *testing.T) {
	cases := []struct {
		lvl  slog.Level
		want string
	}{
		{LevelTrace, "  TRACE"},
		{slog.LevelDebug, "  DEBUG"},
		{slog.LevelInfo, "   INFO"},
		{slog.LevelWarn, "WARNING"},
		{slog.LevelError, "  ERROR"},
		{LevelFatal, "  FATAL"},
		{LevelPanic, "  PANIC"},
	}
	for _, c := range cases {
		buf := &bytes.Buffer{}
		h := NewPrettyHandler(buf, plainPrettyOptions())
		r := slog.NewRecord(time.Now(), c.lvl, "msg", 0)
		require.NoError(t, h.Handle(context.Background(), r))
		require.Contains(t, buf.String(), c.want, "level %v", c.lvl)
	}
}

// TestPrettyHandlerNonCanonicalLevelFallsBack proves levels between the
// canonical seven still render a label rather than an empty column.
func TestPrettyHandlerNonCanonicalLevelFallsBack(t *testing.T) {
	buf := &bytes.Buffer{}
	h := NewPrettyHandler(buf, plainPrettyOptions())
	r := slog.NewRecord(time.Now(), slog.LevelDebug+1, "msg", 0)
	require.NoError(t, h.Handle(context.Background(), r))
	require.Contains(t, buf.String(), "DEBUG+1")
}

func TestPrettyHandlerRelativeTimeUsesRecordTime(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := plainPrettyOptions()
	h := NewPrettyHandler(buf, opts)
	r := slog.NewRecord(opts.StartTimestamp.Add(12345*time.Millisecond), slog.LevelInfo, "msg", 0)
	require.NoError(t, h.Handle(context.Background(), r))
	require.Contains(t, buf.String(), "[  12.345]")
}

func TestPrettyHandlerAbsoluteTime(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := plainPrettyOptions()
	opts.AbsoluteTime = true
	h := NewPrettyHandler(buf, opts)
	at := time.Date(2026, 6, 9, 10, 11, 12, int(13*time.Millisecond), time.UTC)
	r := slog.NewRecord(at, slog.LevelInfo, "msg", 0)
	require.NoError(t, h.Handle(context.Background(), r))
	require.Contains(t, buf.String(), "[2026-06-09 10:11:12.013]")
}

// TestPrettyHandlerResolvesCallerAndTrimsPrefix proves the caller is decoded
// from the record's PC and the configured TrimPrefix is removed from it.
func TestPrettyHandlerResolvesCallerAndTrimsPrefix(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := plainPrettyOptions()
	opts.TrimPrefix = "github.com/openziti/"
	h := NewPrettyHandler(buf, opts)

	slog.New(h).Info("hello")

	line := buf.String()
	require.NotContains(t, line, "github.com/openziti/")
	require.Contains(t, line, " foundation/v2/logging.TestPrettyHandlerResolvesCallerAndTrimsPrefix:")
}

// TestPrettyHandlerColorGating proves UseColor=false produces zero escape
// sequences and UseColor=true produces them.
func TestPrettyHandlerColorGating(t *testing.T) {
	for _, useColor := range []bool{false, true} {
		buf := &bytes.Buffer{}
		opts := plainPrettyOptions()
		opts.UseColor = useColor
		h := NewPrettyHandler(buf, opts)
		r := slog.NewRecord(time.Now(), slog.LevelError, "msg", 0)
		r.AddAttrs(slog.String("k", "v"))
		require.NoError(t, h.Handle(context.Background(), r))
		if useColor {
			require.Contains(t, buf.String(), "\033[", "color enabled must emit ANSI")
		} else {
			require.NotContains(t, buf.String(), "\033[", "color disabled must emit no ANSI")
		}
	}
}

// TestPrettyHandlerDefaultColorFromEnv proves DefaultPrettyOptions leaves color
// off unless PFXLOG_USE_COLOR opts in.
func TestPrettyHandlerDefaultColorFromEnv(t *testing.T) {
	t.Setenv("PFXLOG_USE_COLOR", "")
	require.False(t, DefaultPrettyOptions().UseColor)
	t.Setenv("PFXLOG_USE_COLOR", "true")
	require.True(t, DefaultPrettyOptions().UseColor)
	t.Setenv("PFXLOG_USE_COLOR", "not-a-bool")
	require.False(t, DefaultPrettyOptions().UseColor)
}

// TestPrettyHandlerFieldsSortedAndFormatted checks the pfxlog fields block:
// sorted keys, k=[v] entries, placed before the message.
func TestPrettyHandlerFieldsSortedAndFormatted(t *testing.T) {
	buf := &bytes.Buffer{}
	h := NewPrettyHandler(buf, plainPrettyOptions())
	r := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
	r.AddAttrs(slog.String("b", "2"), slog.String("a", "1"))
	require.NoError(t, h.Handle(context.Background(), r))

	line := buf.String()
	require.Contains(t, line, "{a=[1] b=[2]}")
	require.Less(t, strings.Index(line, "{a=[1]"), strings.Index(line, "msg"), "fields render before the message")
}

// TestPrettyHandlerChannelsAndContext covers the pfxlog _channels/_context
// conventions: rendered after the function as |a, b| and [ctx], excluded
// from the fields block.
func TestPrettyHandlerChannelsAndContext(t *testing.T) {
	buf := &bytes.Buffer{}
	h := NewPrettyHandler(buf, plainPrettyOptions())
	r := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
	r.AddAttrs(
		slog.Any(ChannelsKey, []string{"policyEval", "serviceEval"}),
		slog.String(ContextKey, "ch{edge}"),
		slog.String("k", "v"),
	)
	require.NoError(t, h.Handle(context.Background(), r))

	line := buf.String()
	require.Contains(t, line, "|policyEval, serviceEval|")
	require.Contains(t, line, "[ch{edge}]")
	require.NotContains(t, line, ChannelsKey)
	require.NotContains(t, line, ContextKey)
	require.Contains(t, line, "{k=[v]}")
}

// TestPrettyHandlerFuncAttrFallback proves records without a PC render the
// func attr as the caller and suppress the func/file keys from the fields
// block.
func TestPrettyHandlerFuncAttrFallback(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := plainPrettyOptions()
	opts.TrimPrefix = "github.com/openziti/"
	h := NewPrettyHandler(buf, opts)
	r := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
	r.AddAttrs(
		slog.String("func", "github.com/openziti/ziti/controller.Run"),
		slog.String("file", "controller.go:42"),
	)
	require.NoError(t, h.Handle(context.Background(), r))

	line := buf.String()
	require.Contains(t, line, " ziti/controller.Run:")
	require.NotContains(t, line, "func=[")
	require.NotContains(t, line, "file=[")
}

// TestPrettyHandlerWithAttrsAccumulates proves chained WithAttrs calls append
// rather than replace, and that record attrs override handler attrs with the
// same key.
func TestPrettyHandlerWithAttrsAccumulates(t *testing.T) {
	buf := &bytes.Buffer{}
	base := NewPrettyHandler(buf, plainPrettyOptions())
	h := base.WithAttrs([]slog.Attr{slog.String("a", "1")}).
		WithAttrs([]slog.Attr{slog.String("b", "2")})

	r := slog.NewRecord(time.Now(), slog.LevelInfo, "msg", 0)
	r.AddAttrs(slog.String("b", "3"))
	require.NoError(t, h.Handle(context.Background(), r))

	line := buf.String()
	require.Contains(t, line, "a=[1]", "first WithAttrs batch must survive the second")
	require.Contains(t, line, "b=[3]", "record attr must override handler attr")
	require.NotContains(t, line, "b=[2]")
}
