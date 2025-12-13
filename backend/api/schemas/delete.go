package schemas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type deleteBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}

func (c *Controller) apiDelete(ctx *fiber.Ctx) (err error) {
	var body deleteBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(requestCtx, user, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.db.Transaction(func(tx *gorm.DB) error {
		err = c.schemas.Delete(requestCtx,
			repository.IDsIn(body.IDs),
		)
		if err != nil {
			return err
		}

		return c.schemaScreenshot.Delete(requestCtx, body.IDs)
	})

	return base.NewSuccessResponse(ctx)
}
