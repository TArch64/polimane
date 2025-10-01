package schemas

import (
	"polimane/backend/services/schemascreenshot"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awss3"
	"polimane/backend/services/awssqs"
	"polimane/backend/views"
)

const schemaIdParam = "schemaId"

type ControllerOptions struct {
	fx.In
	Schemas          repositoryschemas.Client
	S3               awss3.Client
	SQS              awssqs.Client
	Renderer         views.Renderer
	SchemaScreenshot schemascreenshot.Interface
}

type Controller struct {
	schemas          repositoryschemas.Client
	sqs              awssqs.Client
	renderer         views.Renderer
	schemaScreenshot schemascreenshot.Interface
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		schemas:          options.Schemas,
		sqs:              options.SQS,
		renderer:         options.Renderer,
		schemaScreenshot: options.SchemaScreenshot,
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
			group.Get("preview", c.apiPreview)
		})
	})
}
