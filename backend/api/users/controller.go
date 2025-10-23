package users

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

const factorIdParam = "factorId"

type ControllerOptions struct {
	fx.In
	WorkosClient workos.Client
	Users        *repositoryusers.Client
	Signals      *signal.Container
}

type Controller struct {
	workosClient workos.Client
	users        *repositoryusers.Client
	signals      *signal.Container
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		workosClient: options.WorkosClient,
		users:        options.Users,
		signals:      options.Signals,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "users/current", func(group fiber.Router) {
		group.Get("", c.apiCurrent)
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
