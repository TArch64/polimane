package sentry

import (
	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"

	"polimane/backend/base"
	"polimane/backend/env"
)

type Container struct {
	Handler fiber.Handler
}

func Provider(environment *env.Environment) (*Container, error) {
	config := environment.Sentry
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
