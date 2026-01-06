package logstdout

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func Provider() *Logger {
	logger := slog.New(newHandler(os.Stdout))
	slog.SetDefault(logger)
	return &Logger{Logger: logger}
}
