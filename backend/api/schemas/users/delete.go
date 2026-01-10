package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	userID, err := base.GetParamID(ctx, ParamUserID)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == userID {
		return base.InvalidRequestErr
	}

	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(reqCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.userSchemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			err = c.userSchemas.DeleteTx(reqCtx, tx,
				repository.HardDelete,
				repository.UserIDEq(userID),
				repository.SchemaIDsIn(body.IDs),
			)

			if err != nil {
				return err
			}

			return c.subscriptionCounters.SchemasCreated.RemoveTx(reqCtx, tx, uint16(len(body.IDs)), userID)
		})

	if err != nil {
		return err
	}

	base.SetResponseUserCounters(ctx, currentUser.Subscription)
	return base.NewSuccessResponse(ctx)
}
