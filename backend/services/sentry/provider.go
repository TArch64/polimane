package sentry

import (
	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
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
		EnableTracing:    true,
		TracesSampleRate: 0.5,
		Environment:      "production",
		SendDefaultPII:   false,
	})

	if err != nil {
		return nil, base.TagError("sentry", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
		),
	)

	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())

	return &Container{IsInitialized: true}, nil
}
