//go:build !dev

package logstdout

import (
	"io"
	"log/slog"
)

func newHandler(writer io.Writer) slog.Handler {
	return slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: true,
	})
}
