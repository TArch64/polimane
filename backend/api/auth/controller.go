package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	"polimane/backend/env"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

const GroupPrefix = "auth"

type ControllerOptions struct {
	fx.In
	Workos  *workos.Client
	Env     *env.Environment
	Users   *repositoryusers.Client
	Signals *signal.Container
}

type Controller struct {
	workos  *workos.Client
	env     *env.Environment
	users   *repositoryusers.Client
	signals *signal.Container
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		workos:  options.Workos,
		env:     options.Env,
		users:   options.Users,
		signals: options.Signals,
	}
}

func (c *Controller) Public(group fiber.Router) {
	base.WithGroup(group, GroupPrefix+"/login", func(group fiber.Router) {
		group.Get("", c.Login)
		group.Get("complete", c.LoginComplete)
	})
}

func (c *Controller) Private(group fiber.Router) {
	group = group.Group(GroupPrefix)
	group.Post("logout", c.Logout)
}
