package logpersistent

import (
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"go.uber.org/fx"

	"polimane/backend/services/appcontext"
	"polimane/backend/services/logstdout"
)

type Logger struct {
	*slog.Logger
}

type ProviderOptions struct {
	fx.In
	Ctx        *appcontext.Ctx
	Cloudwatch *cloudwatchlogs.Client
	Lifecycle  fx.Lifecycle
	Stdout     *logstdout.Logger
}

func Provider(options ProviderOptions) *Logger {
	writer := newWriter(options)
	logger := slog.New(newHandler(writer))

	options.Lifecycle.Append(fx.StopHook(func() error {
		return writer.Close()
	}))

	return &Logger{Logger: logger}
}
