package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

type updateBody struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func (c *Controller) apiUpdate(ctx *fiber.Ctx) (err error) {
	var body updateBody
	if err = ctx.BodyParser(&body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	_, err = c.workosClient.UserManagement.UpdateUser(ctx.Context(), usermanagement.UpdateUserOpts{
		User:      user.WorkosID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})

	if err != nil {
		return err
	}

	c.signals.InvalidateWorkosUserCache.Emit(ctx.Context(), user.WorkosID)
	return base.NewSuccessResponse(ctx)
}
