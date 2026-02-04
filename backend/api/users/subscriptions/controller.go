package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	"polimane/backend/services/subscriptionupdate"
)

type ControllerOptions struct {
	fx.In
	SubscriptionUpdate *subscriptionupdate.Service
}

type Controller struct {
	subscriptionUpdate *subscriptionupdate.Service
}

func Provider(options ControllerOptions) *Controller {
	return &Controller{
		subscriptionUpdate: options.SubscriptionUpdate,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "subscription", func(group fiber.Router) {
		group.Get("plans", c.Plans)
		group.Post("change", c.Change)
	})
}
