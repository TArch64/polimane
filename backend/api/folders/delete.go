package folders

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type DeleteBody struct {
	DeleteSchemas bool `json:"deleteSchemas"`
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	folderID, err := base.GetParamID(ctx, ParamFolderID)
	if err != nil {
		return err
	}

	var body DeleteBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	currentUser := auth.GetSessionUser(ctx)

	folder, err := c.folders.Get(reqCtx,
		repository.IDEq(folderID),
		repository.UserIDEq(currentUser.ID),
	)
	if err != nil {
		return err
	}

	if body.DeleteSchemas {
		err = c.db.WithContext(reqCtx).Transaction(func(tx *gorm.DB) error {
			if err = c.deleteScreenshots(reqCtx, folder); err != nil {
				return err
			}

			return c.deleteFolder(reqCtx, tx, folder, currentUser)
		})
	} else {
		err = c.deleteFolder(reqCtx, c.db, folder, currentUser)
	}

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) deleteFolder(ctx context.Context, tx *gorm.DB, folder *model.Folder, user *model.User) error {
	return c.folders.DeleteTx(ctx, tx,
		repository.IDEq(folder.ID),
		repository.UserIDEq(user.ID),
	)
}

func (c *Controller) deleteScreenshots(ctx context.Context, folder *model.Folder) error {
	var schemaIDs []model.ID
	err := c.userSchemas.ListOut(ctx, &schemaIDs,
		repository.Select("schema_id"),
		repository.Where("folder_id = ?", folder.ID),
	)
	if err != nil {
		return err
	}

	err = c.schemas.Delete(ctx, repository.IDsIn(schemaIDs))
	if err != nil {
		return err
	}

	return c.schemaScreenshot.Delete(ctx, schemaIDs)
}
