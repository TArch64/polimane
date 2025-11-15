package folders

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	repositoryfolders "polimane/backend/repository/folders"
)

type Controller struct {
	folders *repositoryfolders.Client
	db      *gorm.DB
}

type ProviderOptions struct {
	fx.In
	Folders *repositoryfolders.Client
	DB      *gorm.DB
}

func Provider(options ProviderOptions) base.Controller {
	return &Controller{
		folders: options.Folders,
		db:      options.DB,
	}
}

func (c *Controller) Public(_ fiber.Router) {}

func (c *Controller) Private(group fiber.Router) {
	base.WithGroup(group, "folders", func(group fiber.Router) {
		group.Post("", c.apiCreate)
	})
}
