package base

import "github.com/gofiber/fiber/v2"

func WithGroup(group fiber.Router, path string, handler func(group fiber.Router)) {
	handler(group.Group(path))
}
