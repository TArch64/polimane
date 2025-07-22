package base

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Private(group fiber.Router)
	Public(group fiber.Router)
}
