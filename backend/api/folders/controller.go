package folders

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryfolderschemas "polimane/backend/repository/folderschemas"
)

type Controller struct {
	db            *gorm.DB
	folders       *repositoryfolders.Client
	folderSchemas *repositoryfolderschemas.Client
}

type ProviderOptions struct {
	fx.In
	DB            *gorm.DB
	Folders       *repositoryfolders.Client
	FolderSchemas *repositoryfolderschemas.Client
}

func Provider(options ProviderOptions) base.Controller {
	return &Controller{
		db:            options.DB,
		folders:       options.Folders,
		folderSchemas: options.FolderSchemas,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "folders", func(group fiber.Router) {
		group.Post("", c.apiCreate)
	})
}
