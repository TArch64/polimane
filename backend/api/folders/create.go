package folders

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	apischemas "polimane/backend/api/schemas"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemas "polimane/backend/repository/schemas"
	dberror "polimane/backend/services/db/error"
)

type CreateBody struct {
	AddBody
	Name   string `json:"name" validate:"required,min=1,max=255"`
	AsList *bool  `json:"asList"`
}

func (c *Controller) Create(ctx *fiber.Ctx) (err error) {
	var body CreateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	currentUser := auth.GetSessionUser(ctx)
	err = c.userSchemas.FilterByAccess(reqCtx, currentUser, &body.SchemaIDs, model.AccessRead)
	if err != nil {
		return err
	}

	folder := &model.Folder{
		Name:   body.Name,
		UserID: currentUser.ID,
	}

	err = c.db.WithContext(reqCtx).Transaction(func(tx *gorm.DB) error {
		if err = c.folders.InsertTx(reqCtx, tx, folder); err != nil {
			return err
		}

		return c.userSchemas.UpdateTx(reqCtx, tx,
			model.UserSchema{FolderID: &folder.ID},
			repository.SchemaIDsIn(body.SchemaIDs),
		)
	})

	if errors.Is(err, dberror.UniqueConstraintErr) {
		return NameAlreadyInUseErr
	}
	if err != nil {
		return err
	}

	if body.AsList == nil || !*body.AsList {
		return ctx.JSON(folder)
	}

	listFolder, err := c.asListFolder(reqCtx, folder, body.SchemaIDs)
	if err != nil {
		return err
	}

	return ctx.JSON(listFolder)
}

func (c *Controller) asListFolder(ctx context.Context, folder *model.Folder, schemaIDs []model.ID) (*apischemas.ListFolder, error) {
	schema, err := c.schemas.Get(ctx,
		repository.Select("id", "background_color", "screenshoted_at"),
		repository.IDsIn(schemaIDs),
		repositoryschemas.FilterScreenshoted,
		repository.Order("created_at"),
		repository.First,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		schema = nil
		err = nil
	}
	if err != nil {
		return nil, err
	}

	listFolder := &apischemas.ListFolder{
		ID:   folder.ID,
		Name: folder.Name,
	}

	if schema != nil {
		listFolder.BackgroundColor = schema.BackgroundColor
		listFolder.ScreenshotID = &schema.ID
		listFolder.ScreenshotedAt = schema.ScreenshotedAt

		if err = listFolder.AfterScan(); err != nil {
			return nil, err
		}
	}

	return listFolder, nil
}
