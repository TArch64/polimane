package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type deleteBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
	var err error
	var body deleteBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	body.IDs, err = c.userSchemas.FilterByAccess(requestCtx, user.ID, body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.schemas.DeleteMany(ctx.Context(), &repositoryschemas.DeleteOptions{
		SchemaIDs: body.IDs,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
