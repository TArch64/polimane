package schemas

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	"polimane/backend/api/schemas/users"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/views"
)

const schemaIDParam = "schemaID"

type ControllerOptions struct {
	fx.In
	Schemas          *repositoryschemas.Client
	SQS              *awssqs.Client
	Renderer         *views.Renderer
	SchemaScreenshot *schemascreenshot.Service
	UsersController  *users.Controller
}

type Controller struct {
	schemas          *repositoryschemas.Client
	sqs              *awssqs.Client
	renderer         *views.Renderer
	schemaScreenshot *schemascreenshot.Service
	usersController  *users.Controller
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		schemas:          options.Schemas,
		sqs:              options.SQS,
		renderer:         options.Renderer,
		schemaScreenshot: options.SchemaScreenshot,
		usersController:  options.UsersController,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "schemas", func(group fiber.Router) {
		group.Get("", c.apiList)
		group.Post("", c.apiCreate)

		base.WithGroup(group, ":"+schemaIDParam, func(group fiber.Router) {
			group.Get("", c.apiByID)
			group.Delete("", c.apiDelete)
			group.Patch("", c.apiUpdate)
			group.Post("copy", c.apiCopy)
			group.Get("preview", c.apiPreview)

			c.usersController.Private(group)
		})
	})
}
