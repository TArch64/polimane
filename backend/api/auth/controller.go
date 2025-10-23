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

const groupPrefix = "auth"

type ControllerOptions struct {
	fx.In
	WorkosClient workos.Client
	Env          *env.Environment
	Users        *repositoryusers.Client
	Signals      *signal.Container
}

type Controller struct {
	workosClient workos.Client
	env          *env.Environment
	users        *repositoryusers.Client
	signals      *signal.Container
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		workosClient: options.WorkosClient,
		env:          options.Env,
		users:        options.Users,
		signals:      options.Signals,
	}
}

func (c *Controller) Public(group fiber.Router) {
	base.WithGroup(group, groupPrefix+"/login", func(group fiber.Router) {
		group.Get("", c.apiLogin)
		group.Get("complete", c.apiLoginComplete)
	})
}

func (c *Controller) Private(group fiber.Router) {
	group = group.Group(groupPrefix)
	group.Post("logout", c.apiLogout)
}
