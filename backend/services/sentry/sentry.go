package sentry

import (
	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"

	"polimane/backend/env"
)

func Init() (fiber.Handler, error) {
	config := env.Env().Sentry
	if len(config.Dsn) == 0 {
		return nil, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Dsn,
		Release:          config.Release,
		AttachStacktrace: true,
	})

	if err != nil {
		return nil, err
	}

	sentryHandler := sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: true,
	})

	return sentryHandler, nil
}
