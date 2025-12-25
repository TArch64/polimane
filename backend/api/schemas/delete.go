package schemas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type DeleteBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}

func (c *Controller) Delete(ctx *fiber.Ctx) (err error) {
	var body DeleteBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	reqCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(reqCtx, user, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.db.WithContext(reqCtx).Transaction(func(tx *gorm.DB) error {
		err = c.schemas.DeleteTx(reqCtx, tx,
			repository.IDsIn(body.IDs),
		)
		if err != nil {
			return err
		}

		return c.userSchemas.DeleteTx(reqCtx, tx,
			repository.SchemaIDsIn(body.IDs),
		)
	})

	return base.NewSuccessResponse(ctx)
}
