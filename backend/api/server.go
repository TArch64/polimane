package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/api/ping"
	"polimane/backend/api/schemas"
	"polimane/backend/api/users"
	"polimane/backend/env"
	"polimane/backend/services/sentry"
)

type Options struct {
	Protocol  string
	Configure func(config *fiber.Config)
}

func New(options *Options) (*fiber.App, error) {
	sentryHandler, err := sentry.Init()
	if err != nil {
		return nil, err
	}

	config := fiber.Config{
		AppName:      "Polimane",
		ErrorHandler: base.ErrorHandler,
	}

	options.Configure(&config)
	app := fiber.New(config)

	if sentryHandler != nil {
		app.Use(sentryHandler)
	}

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("%s://%s", options.Protocol, env.Env().AppDomain),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-CSRF-Token, Cookie",
		AllowMethods:     "*",
		ExposeHeaders:    "*",
		AllowCredentials: true,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: env.Env().SecretKey,
	}))

	base.InitValidator()

	group := app.Group("/api")
	auth.Group(group)
	ping.Group(group)

	group.Use(auth.NewMiddleware())
	users.Group(group)
	schemas.Group(group)

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Unhandled route:", c.Path())

		return c.
			Status(404).
			JSON(fiber.Map{"error": "Not Found"})
	})

	return app, nil
}
