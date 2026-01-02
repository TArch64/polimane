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

func Provider(options ProviderOptions) (*Logger, error) {
	writer := newWriter(options)

	options.Lifecycle.Append(fx.StopHook(func() error {
		return writer.Close()
	}))

	logger := slog.New(slog.NewJSONHandler(writer, nil))
	return &Logger{Logger: logger}, nil
}
