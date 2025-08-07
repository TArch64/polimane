package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

type updateBody struct {
	FirstName string `json:"firstName" validate:"omitempty,min=1"`
	LastName  string `json:"lastName" validate:"omitempty,min=1"`
	Email     string `json:"email" validate:"omitempty,email"`
}

func (c *Controller) apiUpdate(ctx *fiber.Ctx) (err error) {
	var body updateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if body.FirstName == "" && body.LastName == "" && body.Email == "" {
		return fiber.ErrUnprocessableEntity
	}

	user := auth.GetSessionUser(ctx)

	if err = c.updateUser(ctx.Context(), user.WorkosID, &body); err != nil {
		return err
	}

	if body.Email != "" {
		if err = c.sendEmailVerification(ctx.Context(), user.WorkosID); err != nil {
			return err
		}
	}

	c.signals.InvalidateWorkosUserCache.Emit(ctx.Context(), user.WorkosID)
	return base.NewSuccessResponse(ctx)
}

func (c *Controller) updateUser(ctx context.Context, userID string, body *updateBody) error {
	_, err := c.workosClient.UserManagement().UpdateUser(ctx, usermanagement.UpdateUserOpts{
		User:      userID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	})

	return err
}
