package schemas

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awss3"
)

const schemaIdParam = "schemaId"

type ControllerOptions struct {
	fx.In
	Schemas *repositoryschemas.Client
	S3      awss3.Client
}

type Controller struct {
	schemas *repositoryschemas.Client
	s3      awss3.Client
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		schemas: options.Schemas,
		s3:      options.S3,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "schemas", func(group fiber.Router) {
		group.Get("", c.apiList)
		group.Post("", c.apiCreate)

		base.WithGroup(group, ":"+schemaIdParam, func(group fiber.Router) {
			group.Get("", c.apiById)
			group.Delete("", c.apiDelete)
			group.Patch("", c.apiUpdate)
			group.Post("copy", c.apiCopy)
			group.Patch("screenshot", c.apiUpdateScreenshot)
		})
	})
}
