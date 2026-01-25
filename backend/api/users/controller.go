package users

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	apiauthfactors "polimane/backend/api/users/authfactors"
	"polimane/backend/env"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

type ControllerOptions struct {
	fx.In
	Env                   *env.Environment
	Workos                *workos.Client
	Users                 *repositoryusers.Client
	Signals               *signal.Container
	AuthFactorsController *apiauthfactors.Controller
}

type Controller struct {
	env                   *env.Environment
	workos                *workos.Client
	users                 *repositoryusers.Client
	signals               *signal.Container
	authFactorsController *apiauthfactors.Controller
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		env:                   options.Env,
		workos:                options.Workos,
		users:                 options.Users,
		signals:               options.Signals,
		authFactorsController: options.AuthFactorsController,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "users/current", func(group fiber.Router) {
		group.Get("", c.Current)
		group.Patch("", c.Update)
		group.Delete("", c.Delete)

		c.authFactorsController.Private(group)

		base.WithGroup(group, "email/verify", func(group fiber.Router) {
			group.Post("", c.EmailVerify)
			group.Post("retry", c.EmailVerifyRetry)
		})

		group.Post("password/reset", c.PasswordReset)
	})
}
