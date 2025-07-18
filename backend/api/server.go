package api

import (
	"log"

	"polimane/backend/api/ping"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
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
		AllowOrigins:     env.Instance.AppURL().String(),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Refresh-Token, X-Requested-With, X-CSRF-Token, Cookie",
		AllowMethods:     "*",
		ExposeHeaders:    "*",
		AllowCredentials: true,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: env.Instance.SecretKey,
	}))

	base.InitValidator()

	router := app.Group("/api")

	auth.PublicGroup(router)

	router.Use(auth.NewMiddleware())
	auth.Group(router)
	users.Group(router)
	schemas.Group(router)
	ping.Group(router)

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Unhandled route:", c.Path())

		return c.
			Status(404).
			JSON(fiber.Map{"error": "Not Found"})
	})

	return app, nil
}
