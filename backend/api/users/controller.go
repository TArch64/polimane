package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/signal"

	"polimane/backend/api/base"
	"polimane/backend/services/workos"
)

type Controller struct {
	workosClient *workos.Client
	signals      *signal.Container
}

func Provider(workosClient *workos.Client, signals *signal.Container) base.Controller {
	return &Controller{
		workosClient: workosClient,
		signals:      signals,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	group = group.Group("users/current")
	group.Get("", c.apiGet)
	group.Patch("", c.apiUpdate)
}
