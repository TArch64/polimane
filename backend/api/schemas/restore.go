package schemas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/services/subscriptioncounters"
)

func (c *Controller) Restore(ctx *fiber.Ctx) (err error) {
	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	user := auth.GetSessionUser(ctx)

	err = c.schemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			if err = c.schemas.RestoreTx(reqCtx, tx, body.IDs); err != nil {
				return err
			}

			return c.subscriptionCounters.SchemasCreated.ChangeTx(reqCtx, tx, subscriptioncounters.ChangeSet{
				user.ID: int16(len(body.IDs)),
			})
		})

	if err != nil {
		return err
	}

	base.SetResponseUserCounters(ctx, user.Subscription)
	return base.NewSuccessResponse(ctx)
}
