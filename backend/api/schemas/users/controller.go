package users

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

const schemaIdParam = "schemaId"
const userIdParam = "userId"

type ControllerOptions struct {
	fx.In
	Users       *repositoryusers.Client
	UserSchemas *repositoryuserschemas.Client
}

type Controller struct {
	users       *repositoryusers.Client
	userSchemas *repositoryuserschemas.Client
}

func Provider(options ControllerOptions) *Controller {
	return &Controller{
		users:       options.Users,
		userSchemas: options.UserSchemas,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "users", func(group fiber.Router) {
		group.Get("", c.apiList)
		group.Post("", c.apiAdd)

		base.WithGroup(group, ":"+userIdParam, func(group fiber.Router) {
			group.Delete("", c.apiDelete)
		})
	})
}
