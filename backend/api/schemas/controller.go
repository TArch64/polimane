package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

type Controller struct {
	schemas *repositoryschemas.Client
}

func Provider(schemas *repositoryschemas.Client) base.Controller {
	return &Controller{
		schemas: schemas,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	group = group.Group("schemas")
	group.Get("", c.apiList)
	group.Post("", c.apiCreate)

	group = group.Group(":schemaId")
	group.Get("", c.apiById)
	group.Delete("", c.apiDelete)
	group.Patch("", c.apiUpdate)
	group.Post("copy", c.apiCopy)
}
