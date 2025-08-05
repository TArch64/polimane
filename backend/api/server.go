package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"go.uber.org/fx"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/services/sentry"
)

type Options struct {
	Protocol  string
	Configure func(config *fiber.Config)
}

type ServerOptions struct {
	fx.In
	Controllers    []base.Controller `group:"controllers"`
	Options        *Options
	Sentry         *sentry.Container
	Env            *env.Environment
	AuthMiddleware *auth.Middleware
}

func Provider(options ServerOptions) (*fiber.App, error) {
	config := fiber.Config{
		AppName:      "Polimane",
		ErrorHandler: base.ErrorHandler,
	}

	options.Options.Configure(&config)
	app := fiber.New(config)

	if options.Sentry.Handler != nil {
		app.Use(options.Sentry.Handler)
	}

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     options.Env.AppURL().String(),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Refresh-Token, X-Requested-With, X-CSRF-Token, Cookie",
		AllowMethods:     "*",
		ExposeHeaders:    "*",
		AllowCredentials: true,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: options.Env.SecretKey,
	}))

	base.InitValidator()

	base.WithGroup(app, "/api", func(group fiber.Router) {
		for _, controller := range options.Controllers {
			controller.Public(group)
		}

		group.Use(options.AuthMiddleware.Handler)

		for _, controller := range options.Controllers {
			controller.Private(group)
		}
	})

	app.Use(apiNotFound)
	return app, nil
}
