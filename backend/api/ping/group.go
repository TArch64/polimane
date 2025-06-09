package ping

import "github.com/gofiber/fiber/v2"

func Group(group fiber.Router) {
	group = group.Group("ping")
	group.Get("", apiPing)
}
