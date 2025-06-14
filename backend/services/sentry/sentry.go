package sentry

import (
	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"

	"polimane/backend/env"
)

func Init() (fiber.Handler, error) {
	dsn := env.Env().Sentry.Dsn
	if len(dsn) == 0 {
		return nil, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
		Debug:            true,
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
