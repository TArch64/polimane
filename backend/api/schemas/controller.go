package schemas

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/views"
)

const schemaIdParam = "schemaId"

type ControllerOptions struct {
	fx.In
	Schemas          *repositoryschemas.Client
	UserSchemas      *repositoryuserschemas.Client
	SQS              *awssqs.Client
	Renderer         *views.Renderer
	SchemaScreenshot *schemascreenshot.Service
}

type Controller struct {
	schemas          *repositoryschemas.Client
	userSchemas      *repositoryuserschemas.Client
	sqs              *awssqs.Client
	renderer         *views.Renderer
	schemaScreenshot *schemascreenshot.Service
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		schemas:          options.Schemas,
		userSchemas:      options.UserSchemas,
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
			group.Get("users", c.apiUsers)
		})
	})
}
