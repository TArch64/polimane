package folders

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryfolderschemas "polimane/backend/repository/folderschemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

const folderIDParam = "folderID"

type Controller struct {
	db            *gorm.DB
	folders       *repositoryfolders.Client
	folderSchemas *repositoryfolderschemas.Client
	userSchemas   *repositoryuserschemas.Client
}

type ProviderOptions struct {
	fx.In
	DB            *gorm.DB
	Folders       *repositoryfolders.Client
	FolderSchemas *repositoryfolderschemas.Client
	UserSchemas   *repositoryuserschemas.Client
}

func Provider(options ProviderOptions) base.Controller {
	return &Controller{
		db:            options.DB,
		folders:       options.Folders,
		folderSchemas: options.FolderSchemas,
		userSchemas:   options.UserSchemas,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "folders", func(group fiber.Router) {
		group.Post("", c.apiCreate)

		base.WithGroup(group, ":"+folderIDParam, func(group fiber.Router) {
			group.Post("schemas", c.apiAddSchema)
		})
	})
}
