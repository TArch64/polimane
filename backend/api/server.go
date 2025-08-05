package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/services/sentry"
)

type Options struct {
	Protocol  string
	Configure func(config *fiber.Config)
}

func Provider(
	controllers []base.Controller, // group:"controllers"
	options *Options,
	sentry *sentry.Container,
	environment *env.Environment,
	authMiddleware *auth.Middleware,
) (*fiber.App, error) {
	config := fiber.Config{
		AppName:      "Polimane",
		ErrorHandler: base.ErrorHandler,
	}

	options.Configure(&config)
	app := fiber.New(config)

	if sentry.Handler != nil {
		app.Use(sentry.Handler)
	}

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     environment.AppURL().String(),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Refresh-Token, X-Requested-With, X-CSRF-Token, Cookie",
		AllowMethods:     "*",
		ExposeHeaders:    "*",
		AllowCredentials: true,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: environment.SecretKey,
	}))

	base.InitValidator()

	base.WithGroup(app, "/api", func(group fiber.Router) {
		for _, controller := range controllers {
			controller.Public(group)
		}

		group.Use(authMiddleware.Handler)

		for _, controller := range controllers {
			controller.Private(group)
		}
	})

	app.Use(apiNotFound)
	return app, nil
}
