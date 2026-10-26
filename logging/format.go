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
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
)

// ReplaceAttr is a slog.HandlerOptions.ReplaceAttr callback that coerces
// slog's built-in JSON and text output into the pfxlog shape: the "level"
// value becomes the canonical lowercase name (see LevelName), and the nested
// "source" attr is dropped. Only top-level attrs are modified. Callers that
// want flat "file"/"func" keys use NewJSONHandler or NewTextHandler, which pair
// this with caller flattening.
func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	if len(groups) > 0 {
		return a
	}
	switch a.Key {
	case slog.LevelKey:
		if lvl, ok := a.Value.Any().(slog.Level); ok {
			return slog.String(slog.LevelKey, LevelName(lvl))
		}
	case slog.SourceKey:
		return slog.Attr{}
	}
	return a
}

// NewJSONHandler returns a slog.Handler that writes pfxlog-shaped JSON to out:
// lowercase canonical level names, and the caller as flat "file" and "func"
// keys rather than a nested "source" object. Records with a zero PC carry no
// caller keys of their own, so a bridge that resolves the caller elsewhere can
// supply "file"/"func" as attrs. Writes are synchronous; wrap the result in an
// AsyncHandler to move them off the caller's goroutine. The handler does no
// level gating. out defaults to os.Stderr when nil.
func NewJSONHandler(out io.Writer) slog.Handler {
	if out == nil {
		out = os.Stderr
	}
	return &sourceFlattener{parent: slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level:       LevelTrace,
		ReplaceAttr: ReplaceAttr,
	})}
}

// NewTextHandler is NewJSONHandler with slog's key=value text output in place
// of JSON (level=info msg=... file=... func=...).
func NewTextHandler(out io.Writer) slog.Handler {
	if out == nil {
		out = os.Stderr
	}
	return &sourceFlattener{parent: slog.NewTextHandler(out, &slog.HandlerOptions{
		Level:       LevelTrace,
		ReplaceAttr: ReplaceAttr,
	})}
}

// sourceFlattener decodes a record's PC into flat "file" and "func" attrs
// ahead of the record's own attrs. Records with a zero PC pass through
// unchanged.
type sourceFlattener struct {
	parent slog.Handler
}

var _ slog.Handler = (*sourceFlattener)(nil)

func (h *sourceFlattener) Enabled(ctx context.Context, level slog.Level) bool {
	return h.parent.Enabled(ctx, level)
}

func (h *sourceFlattener) Handle(ctx context.Context, r slog.Record) error {
	if r.PC == 0 {
		return h.parent.Handle(ctx, r)
	}
	frames := runtime.CallersFrames([]uintptr{r.PC})
	frame, _ := frames.Next()
	r2 := slog.NewRecord(r.Time, r.Level, r.Message, 0)
	r2.AddAttrs(
		slog.String("file", fmt.Sprintf("%s:%d", frame.File, frame.Line)),
		slog.String("func", frame.Function),
	)
	r.Attrs(func(a slog.Attr) bool {
		r2.AddAttrs(a)
		return true
	})
	return h.parent.Handle(ctx, r2)
}

func (h *sourceFlattener) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &sourceFlattener{parent: h.parent.WithAttrs(attrs)}
}

func (h *sourceFlattener) WithGroup(name string) slog.Handler {
	return &sourceFlattener{parent: h.parent.WithGroup(name)}
}
