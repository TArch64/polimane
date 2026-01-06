package schemas

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/api/schemas/users"
	"polimane/backend/model"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/schemadelete"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/services/subscriptioncounters"
	"polimane/backend/views"
)

const ParamSchemaID = "schemaID"
const ParamDefSchemaID = ":" + ParamSchemaID

type ControllerOptions struct {
	fx.In
	Schemas              *repositoryschemas.Client
	Folders              *repositoryfolders.Client
	UserSchemas          *repositoryuserschemas.Client
	SQS                  *awssqs.Client
	S3                   *s3.Client
	Renderer             *views.Renderer
	SchemaScreenshot     *schemascreenshot.Service
	SchemaDelete         *schemadelete.Service
	UsersController      *users.Controller
	SubscriptionCounters *subscriptioncounters.Service
}

type Controller struct {
	schemas              *repositoryschemas.Client
	folders              *repositoryfolders.Client
	userSchemas          *repositoryuserschemas.Client
	sqs                  *awssqs.Client
	renderer             *views.Renderer
	schemaScreenshot     *schemascreenshot.Service
	schemaDelete         *schemadelete.Service
	usersController      *users.Controller
	subscriptionCounters *subscriptioncounters.Service
}

func Provider(options ControllerOptions) base.Controller {
	return &Controller{
		schemas:              options.Schemas,
		folders:              options.Folders,
		userSchemas:          options.UserSchemas,
		sqs:                  options.SQS,
		renderer:             options.Renderer,
		schemaScreenshot:     options.SchemaScreenshot,
		schemaDelete:         options.SchemaDelete,
		usersController:      options.UsersController,
		subscriptionCounters: options.SubscriptionCounters,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "schemas", func(group fiber.Router) {
		group.Get("", c.List)
		group.Post("", c.Create)
		group.Delete("delete", c.Delete)
		group.Delete("delete-permanently", c.DeletePermanently)
		group.Post("restore", c.Restore)
		group.Get("deleted", c.ListDeleted)

		c.usersController.Private(group)

		base.WithGroup(group, ParamDefSchemaID, func(group fiber.Router) {
			group.Get("", c.ByID)
			group.Patch("", c.Update)
			group.Post("copy", c.Copy)
			group.Get("preview", c.Preview)
		})
	})
}

func (c *Controller) filterSchemaIDsByAccess(ctx *fiber.Ctx, IDs *[]model.ID) error {
	return c.userSchemas.FilterByAccess(ctx.Context(), auth.GetSessionUser(ctx), IDs, model.AccessAdmin)
}
