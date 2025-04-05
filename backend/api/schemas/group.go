package schemas

import "github.com/gofiber/fiber/v2"

func Group(group fiber.Router) {
	group = group.Group("schemas")
}
