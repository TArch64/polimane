package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/api/schemas"
	"polimane/backend/api/users"
	"polimane/backend/env"
)

type Config func(config *fiber.Config)

func New(configFns ...Config) *fiber.App {
	config := fiber.Config{
		AppName:      "Polimane",
		ErrorHandler: base.ErrorHandler,
	}

	for _, configure := range configFns {
		configure(&config)
	}

	app := fiber.New(config)

	app.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
	}))

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: env.Env().FrontendOrigin,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: env.Env().SecretKey,
	}))

	base.InitValidator()

	group := app.Group("/api")
	auth.Group(group)

	group.Use(auth.Middleware)
	users.Group(group)
	schemas.Group(group)

	return app
}
