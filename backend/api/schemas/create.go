package schemas

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/subscriptioncounters"
)

type CreateBody struct {
	Name        string             `json:"name" validate:"required,min=1,max=255"`
	Layout      model.SchemaLayout `json:"layout" validate:"required,oneof=linear radial"`
	FolderIDStr *string            `json:"folderId" validate:"omitempty,uuid"`
}

func (l *CreateBody) FolderID() *model.ID {
	if l.FolderIDStr == nil {
		return nil
	}
	id := model.MustStringToID(*l.FolderIDStr)
	return &id
}

func (c *Controller) Create(ctx *fiber.Ctx) (err error) {
	var body CreateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	if !c.subscriptionCounters.SchemasCreated.CanAdd(user.Subscription, 1) {
		return base.SchemasCreatedLimitReachedErr
	}

	reqCtx := ctx.Context()
	folderID := body.FolderID()
	if folderID != nil {
		err = c.folders.HasAccess(reqCtx, user.ID, *folderID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			folderID = nil
			err = nil
		}
		if err != nil {
			return err
		}
	}

	var schema *model.Schema

	err = c.schemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			schema, err = c.schemas.CreateTx(reqCtx, tx, &repositoryschemas.CreateOptions{
				User:     user,
				Name:     body.Name,
				Layout:   body.Layout,
				FolderID: folderID,
			})

			if err != nil {
				return err
			}

			return c.subscriptionCounters.SchemasCreated.ChangeTx(reqCtx, tx, subscriptioncounters.ChangeSet{
				user.ID: 1,
			})
		})

	if err != nil {
		return err
	}

	base.SetResponseUserCounters(ctx, user.Subscription)

	if err = c.updateScreenshot(reqCtx, schema.ID, false); err != nil {
		return err
	}

	return ctx.JSON(NewListSchema(schema, model.AccessAdmin))
}
