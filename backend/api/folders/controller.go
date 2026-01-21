package folders

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/schemascreenshot"
)

const ParamFolderID = "folderID"
const ParamDefFolderID = ":" + ParamFolderID

var (
	NameAlreadyInUseErr = base.NewReasonedError(fiber.StatusBadRequest, "AlreadyInUse[Name]")
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
			group.Post("schemas", c.AddSchemas)
			group.Delete("schemas", c.RemoveSchemas)
		})
	})
}

func (c *Controller) getFolderID(ctx *fiber.Ctx) (model.ID, error) {
	folderID, err := base.GetParamID(ctx, ParamFolderID)
	if err != nil {
		return folderID, err
	}

	if err = c.checkFolderAccess(ctx, folderID); err != nil {
		return folderID, err
	}

	return folderID, nil
}

func (c *Controller) checkFolderAccess(ctx *fiber.Ctx, folderID model.ID) error {
	user := auth.GetSessionUser(ctx)
	return c.folders.HasAccess(ctx.Context(), user.ID, folderID)
}

func (c *Controller) filterSchemaIDsByAccess(ctx *fiber.Ctx, IDs *[]model.ID) error {
	return c.userSchemas.FilterByAccess(ctx.Context(), auth.GetSessionUser(ctx), IDs, model.AccessRead)
}

type BulkSchemasBody struct {
	SchemaIDs []model.ID `json:"schemaIds" validate:"required,dive,required"`
}
