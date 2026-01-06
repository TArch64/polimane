//go:build dev

package logpersistent

import (
	"io"
	"log/slog"

	"github.com/Marlliton/slogpretty"
)

func newHandler(writer io.Writer) slog.Handler {
	return slogpretty.New(writer, &slogpretty.Options{
		AddSource:  true,
		Colorful:   true,
		Multiline:  true,
		TimeFormat: slogpretty.DefaultTimeFormat,
	})
}
