package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/services/subscriptioncounters"
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

	affectedUsers, err := c.getAffectedUsersOnDelete(reqCtx, body.IDs)
	if err != nil {
		return err
	}

	err = c.schemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			if err = c.schemas.DeleteSoftTx(reqCtx, tx, user, body.IDs); err != nil {
				return err
			}

			return c.subscriptionCounters.SchemasCreated.ChangeTx(reqCtx, tx, affectedUsers)
		})

	if err != nil {
		return err
	}

	base.SetResponseUserCounters(ctx, user.Subscription)
	return base.NewSuccessResponse(ctx)
}

func (c *Controller) getAffectedUsersOnDelete(ctx context.Context, schemaIDs []model.ID) (subscriptioncounters.ChangeSet, error) {
	rows, err := c.userSchemas.ListRows(ctx,
		repository.Table("user_schemas"),
		repository.Select("user_id", "COUNT(schema_id)"),
		repository.SchemaIDsIn(schemaIDs),
		repository.Group("user_id"),
	)

	if err != nil {
		return nil, err
	}

	changeSet := make(subscriptioncounters.ChangeSet)

	for rows.Next() {
		var id model.ID
		var count int16

		if err = rows.Scan(&id, &count); err != nil {
			return nil, err
		}

		changeSet[id] = -count
	}

	return changeSet, err
}
