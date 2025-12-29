package logpersistent

import (
	"log/slog"

	"go.uber.org/fx"
)

type Logger struct {
	*slog.Logger
}

type ProviderOptions struct {
	fx.In
}

func Provider(options ProviderOptions) (*Logger, error) {
	writer, err := newWriter(options)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(writer, nil))
	return &Logger{Logger: logger}, nil
}
