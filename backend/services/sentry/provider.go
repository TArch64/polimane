package sentry

import (
	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/base"
	"polimane/backend/env"
)

type Options struct {
	fx.In
	Env *env.Environment
}

type Container struct {
	Handler fiber.Handler
}

func Provider(options Options) (*Container, error) {
	config := options.Env.Sentry
	if len(config.Dsn) == 0 {
		return &Container{Handler: nil}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Dsn,
		Release:          config.Release,
		AttachStacktrace: true,
	})

	if err != nil {
		return nil, base.TagError("sentry", err)
	}

	sentryHandler := sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: true,
	})

	return &Container{Handler: sentryHandler}, nil
}
