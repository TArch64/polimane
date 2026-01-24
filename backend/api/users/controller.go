package users

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	"polimane/backend/env"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

const ParamFactorID = "factorID"
const ParamDefFactorID = ":" + ParamFactorID

type ControllerOptions struct {
	fx.In
	Env     *env.Environment
	Workos  *workos.Client
	Users   *repositoryusers.Client
	Signals *signal.Container
}

type Controller struct {
	env     *env.Environment
	workos  *workos.Client
	users   *repositoryusers.Client
	signals *signal.Container
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		env:     options.Env,
		workos:  options.Workos,
		users:   options.Users,
		signals: options.Signals,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "users/current", func(group fiber.Router) {
		group.Get("", c.Current)
		group.Patch("", c.Update)
		group.Delete("", c.Delete)
		group.Get("plans", c.Plans)

		base.WithGroup(group, "email/verify", func(group fiber.Router) {
			group.Post("", c.EmailVerify)
			group.Post("retry", c.EmailVerifyRetry)
		})

		group.Post("password/reset", c.PasswordReset)

		base.WithGroup(group, "auth-factors", func(group fiber.Router) {
			group.Get("", c.ListAuthFactors)
			group.Post("", c.AuthFactorCreate)
			group.Post("init", c.AuthFactorsInit)
			group.Delete(ParamDefFactorID, c.AuthFactorDelete)
		})
	})
}
