//go:build dev

package fxlogger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"polimane/backend/services/logstdout"
)

var Provider = fx.WithLogger(func(stdout *logstdout.Logger) fxevent.Logger {
	return &fxevent.SlogLogger{
		Logger: stdout.Logger,
	}
})
