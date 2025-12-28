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
	folderID, err := c.getFolderID(ctx)
	if err != nil {
		return err
	}

	var body DeleteBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	currentUser := auth.GetSessionUser(ctx)

	if body.DeleteSchemas {
		err = c.db.WithContext(reqCtx).Transaction(func(tx *gorm.DB) error {
			if err = c.deleteScreenshots(reqCtx, folderID); err != nil {
				return err
			}

			return c.deleteFolder(reqCtx, tx, folderID, currentUser)
		})
	} else {
		err = c.deleteFolder(reqCtx, c.db, folderID, currentUser)
	}

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) deleteFolder(ctx context.Context, tx *gorm.DB, folderID model.ID, user *model.User) error {
	return c.folders.DeleteTx(ctx, tx,
		repository.IDEq(folderID),
		repository.UserIDEq(user.ID),
	)
}

func (c *Controller) deleteScreenshots(ctx context.Context, folderID model.ID) error {
	var schemaIDs []model.ID
	err := c.userSchemas.ListOut(ctx, &schemaIDs,
		repository.Select("schema_id"),
		repository.Where("folder_id = ?", folderID),
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
