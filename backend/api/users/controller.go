package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
)

type Controller struct{}

func Provider() base.Controller {
	return &Controller{}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	group = group.Group("users")
	group.Get("current", c.apiCurrent)
}
