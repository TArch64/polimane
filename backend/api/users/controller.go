package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
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

	group.Post("email/verify", c.apiEmailVerify)
	group.Post("email/verify/retry", c.apiEmailVerifyRetry)

	group.Post("password/reset", c.apiPasswordReset)

	group.Get("auth-factors", c.apiListAuthFactors)
}
