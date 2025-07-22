package api

import (
	"log"

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

	router := app.Group("/api")

	for _, controller := range controllers {
		controller.Public(router)
	}

	router.Use(authMiddleware.Handler)

	for _, controller := range controllers {
		controller.Private(router)
	}

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Unhandled route:", c.Path())

		return c.
			Status(404).
			JSON(fiber.Map{"error": "Not Found"})
	})

	return app, nil
}
