package folders

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/schemascreenshot"
)

const ParamFolderID = "folderID"
const ParamDefFolderID = ":" + ParamFolderID

var (
	NameAlreadyInUseErr = base.NewReasonedError(fiber.StatusBadRequest, "NameAlreadyInUse")
)

type Controller struct {
	db               *gorm.DB
	folders          *repositoryfolders.Client
	userSchemas      *repositoryuserschemas.Client
	schemas          *repositoryschemas.Client
	schemaScreenshot *schemascreenshot.Service
}

type ProviderOptions struct {
	fx.In
	DB               *gorm.DB
	Folders          *repositoryfolders.Client
	UserSchemas      *repositoryuserschemas.Client
	Schemas          *repositoryschemas.Client
	SchemaScreenshot *schemascreenshot.Service
}

func Provider(options ProviderOptions) base.Controller {
	return &Controller{
		db:               options.DB,
		folders:          options.Folders,
		userSchemas:      options.UserSchemas,
		schemas:          options.Schemas,
		schemaScreenshot: options.SchemaScreenshot,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "folders", func(group fiber.Router) {
		group.Get("", c.List)
		group.Post("", c.Create)

		base.WithGroup(group, ParamDefFolderID, func(group fiber.Router) {
			group.Get("", c.ByID)
			group.Patch("", c.Update)
			group.Delete("", c.Delete)
			group.Post("schemas", c.AddSchema)
		})
	})
}
