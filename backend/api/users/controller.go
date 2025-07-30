package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

const factorIdParam = "factorId"

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
	base.WithGroup(group, "users/current", func(group fiber.Router) {
		group.Get("", c.apiGet)
		group.Patch("", c.apiUpdate)

		base.WithGroup(group, "email/verify", func(group fiber.Router) {
			group.Post("", c.apiEmailVerify)
			group.Post("retry", c.apiEmailVerifyRetry)
		})

		group.Post("password/reset", c.apiPasswordReset)

		base.WithGroup(group, "auth-factors", func(group fiber.Router) {
			group.Get("", c.apiListAuthFactors)
			group.Post("", c.apiAuthFactorCreate)
			group.Post("init", c.apiAuthFactorsInit)
			group.Delete(":"+factorIdParam, c.apiAuthFactorDelete)
		})
	})
}
