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

	app.Use(getErrorHandlerMiddleware())

	app.Use(helmet.New(helmet.Config{
		XSSProtection: "1; mode=block",

		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' https:; connect-src 'self' https://api.workos.com; frame-src 'none';",

		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: false,
		HSTSPreloadEnabled:    true,

		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		ReferrerPolicy:     "no-referrer",

		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "same-origin",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     options.Env.AppURL.String(),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
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

	app.Use(NotFound)
	return app, nil
}
