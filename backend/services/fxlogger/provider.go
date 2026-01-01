package fxlogger

import (
	"go.uber.org/fx/fxevent"

	"polimane/backend/services/logstdout"
)

func Provider(stdout *logstdout.Logger) fxevent.Logger {
	return &fxevent.SlogLogger{
		Logger: stdout.Logger,
	}
}
