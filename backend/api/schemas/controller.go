package schemas

import (
	"gorm.io/gorm"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/base"
	"polimane/backend/api/schemas/users"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/views"
)

const schemaIDParam = "schemaID"

type ControllerOptions struct {
	fx.In
	Schemas          *repositoryschemas.Client
	Folders          *repositoryfolders.Client
	UserSchemas      *repositoryuserschemas.Client
	SQS              *awssqs.Client
	S3               *s3.Client
	DB               *gorm.DB
	Renderer         *views.Renderer
	SchemaScreenshot *schemascreenshot.Service
	UsersController  *users.Controller
}

type Controller struct {
	schemas          *repositoryschemas.Client
	folders          *repositoryfolders.Client
	userSchemas      *repositoryuserschemas.Client
	sqs              *awssqs.Client
	s3               *s3.Client
	db               *gorm.DB
	renderer         *views.Renderer
	schemaScreenshot *schemascreenshot.Service
	usersController  *users.Controller
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		schemas:          options.Schemas,
		folders:          options.Folders,
		userSchemas:      options.UserSchemas,
		sqs:              options.SQS,
		s3:               options.S3,
		db:               options.DB,
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
		group.Delete("delete-many", c.apiDelete)

		c.usersController.Private(group)

		base.WithGroup(group, ":"+schemaIDParam, func(group fiber.Router) {
			group.Get("", c.apiByID)
			group.Patch("", c.apiUpdate)
			group.Post("copy", c.apiCopy)
			group.Get("preview", c.apiPreview)
		})
	})
}
