package slogcar

import (
	"context"
	"io"
	"log/slog"

	"github.com/lingdor/go-logcar/entity"
)

type slogHandler struct {
	baseHandler slog.Handler
	w           io.Writer
}

func (s *slogHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return s.baseHandler.Enabled(ctx, l)
}

func (s *slogHandler) Handle(ctx context.Context, r slog.Record) error {

	carLevel := entity.LogLevelInfo
	switch r.Level {
	case slog.LevelDebug:
		carLevel = entity.LogLevelDebug
	case slog.LevelError:
		carLevel = entity.LogLevelError
	case slog.LevelWarn:
		carLevel = entity.LogLevelWarn
	}
	s.w.Write([]byte{byte(carLevel)})
	return s.baseHandler.Handle(ctx, r)
}

func (s *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {

	s.baseHandler = s.baseHandler.WithAttrs(attrs)
	return s
}

func (s *slogHandler) WithGroup(name string) slog.Handler {

	s.baseHandler = s.baseHandler.WithGroup(name)
	return s
}

func NewJsonHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return &slogHandler{
		baseHandler: slog.NewJSONHandler(w, opts),
		w:           w,
	}
}
func NewFromBaseHandler(handler slog.Handler, w io.Writer) slog.Handler {
	return &slogHandler{
		baseHandler: handler,
		w:           w,
	}
}

func NewTextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return &slogHandler{
		baseHandler: slog.NewTextHandler(w, opts),
		w:           w,
	}
}

// func NewDirectHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
// 	return &slogHandler{
// 		baseHandler: slog.NewTextHandler(w, opts),
// 		w:           &carWriter{w: w},
// 	}
// }
