package schemas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/subscriptioncounters"
)

func (c *Controller) Copy(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, ParamSchemaID)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	if !c.subscriptionCounters.SchemasCreated.CanAdd(user, 1) {
		return base.SchemasCreatedLimitReachedErr
	}

	reqCtx := ctx.Context()
	var schema *model.Schema

	err = c.schemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			schema, err = c.schemas.CopyTx(reqCtx, tx, &repositoryschemas.CopyOptions{
				User:     user,
				SchemaID: schemaID,
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
