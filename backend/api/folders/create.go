package folders

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryfolderschemas "polimane/backend/repository/folderschemas"
)

type createBody struct {
	Name        string     `json:"name" validate:"required,min=1,max=255"`
	SchemaIDs   []model.ID `json:"schemaIds" validate:"dive,required"`
	OldFolderID *model.ID  `json:"oldFolderId"`
}

func (c *Controller) apiCreate(ctx *fiber.Ctx) (err error) {
	var body createBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	folder := &model.Folder{
		Name:   body.Name,
		UserID: auth.GetSessionUser(ctx).ID,
	}

	requestCtx := ctx.Context()
	err = c.db.Transaction(func(tx *gorm.DB) error {
		if err = c.folders.CreateTx(requestCtx, tx, folder); err != nil {
			return err
		}

		return c.folderSchemas.AddManyTx(requestCtx, tx, &repositoryfolderschemas.AddManyOptions{
			FolderID:  folder.ID,
			SchemaIDs: body.SchemaIDs,
		})
	})

	if err != nil {
		return err
	}

	return ctx.JSON(folder)
}
