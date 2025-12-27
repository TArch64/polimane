package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
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

	if err = c.schemas.SoftDelete(reqCtx, body.IDs); err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
