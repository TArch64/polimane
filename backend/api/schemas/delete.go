package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Controller) Delete(ctx *fiber.Ctx) (err error) {
	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	reqCtx := ctx.Context()

	affectedUserIDs, err := c.getAffectedUsersOnDelete(reqCtx, body.IDs)
	if err != nil {
		return err
	}

	err = c.schemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			if err = c.schemas.DeleteSoftTx(reqCtx, tx, user, body.IDs); err != nil {
				return err
			}

			return c.subscriptionCounters.SchemasCreated.RemoveTx(reqCtx, tx, len(body.IDs), affectedUserIDs...)
		})

	if err != nil {
		return err
	}

	base.SetResponseUserCounters(ctx, user.Subscription)
	return base.NewSuccessResponse(ctx)
}

func (c *Controller) getAffectedUsersOnDelete(ctx context.Context, schemaIDs []model.ID) (out []model.ID, err error) {
	err = c.userSchemas.ListOut(ctx, &out,
		repository.Select("DISTINCT ON (user_id) user_id"),
		repository.SchemaIDsIn(schemaIDs),
	)
	return
}
