package auth

import "github.com/gofiber/fiber/v2"

func Group(group fiber.Router) {
	group = group.Group("auth")
	group.Get("login", apiLogin)
	group.Get("login/complete", apiLoginComplete)
}
