package auth

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/env"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

const groupPrefix = "auth"

type Controller struct {
	workosClient *workos.Client
	env          *env.Environment
	users        *repositoryusers.Client
	signals      *signal.Container
}

func Provider(
	workosClient *workos.Client,
	environment *env.Environment,
	users *repositoryusers.Client,
	signals *signal.Container,
) base.Controller {
	return &Controller{
		workosClient: workosClient,
		env:          environment,
		users:        users,
		signals:      signals,
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
