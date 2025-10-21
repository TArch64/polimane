package sentry

import (
	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"

	"polimane/backend/base"
	"polimane/backend/env"
)

type Options struct {
	fx.In
	Env *env.Environment
}

type Container struct {
	IsInitialized bool
}

func Provider(options Options) (*Container, error) {
	config := options.Env.Sentry
	if len(config.Dsn) == 0 {
		return &Container{}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Dsn,
		Release:          config.Release,
		AttachStacktrace: true,
		Environment:      "production",
	})

	if err != nil {
		return nil, base.TagError("sentry", err)
	}

	return &Container{IsInitialized: true}, nil
}
