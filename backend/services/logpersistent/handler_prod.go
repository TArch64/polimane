//go:build !dev

package logpersistent

import (
	"io"
	"log/slog"
)

func newHandler(writer io.Writer) slog.Handler {
	return slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: true,
	})
}
