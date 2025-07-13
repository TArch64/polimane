package auth

import "github.com/gofiber/fiber/v2"

const groupPrefix = "auth"

func Group(group fiber.Router) {
	group = group.Group(groupPrefix)
}

func PublicGroup(group fiber.Router) {
	group = group.Group(groupPrefix)
	group.Get("login", apiLogin)
	group.Get("login/complete", apiLoginComplete)
}
