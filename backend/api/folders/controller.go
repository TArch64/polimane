package folders

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

const folderIDParam = "folderID"

var (
	NameAlreadyInUseErr = base.NewReasonedError(fiber.StatusBadRequest, "NameAlreadyInUse")
)

type Controller struct {
	db          *gorm.DB
	folders     *repositoryfolders.Client
	userSchemas *repositoryuserschemas.Client
	schemas     *repositoryschemas.Client
}

type ProviderOptions struct {
	fx.In
	DB          *gorm.DB
	Folders     *repositoryfolders.Client
	UserSchemas *repositoryuserschemas.Client
	Schemas     *repositoryschemas.Client
}

func Provider(options ProviderOptions) base.Controller {
	return &Controller{
		db:          options.DB,
		folders:     options.Folders,
		userSchemas: options.UserSchemas,
		schemas:     options.Schemas,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "folders", func(group fiber.Router) {
		group.Get("", c.apiList)
		group.Post("", c.apiCreate)

		base.WithGroup(group, ":"+folderIDParam, func(group fiber.Router) {
			group.Get("", c.apiByID)
			group.Patch("", c.apiUpdate)
			group.Delete("", c.apiDelete)
			group.Post("schemas", c.apiAddSchema)
		})
	})
}
