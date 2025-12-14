package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type UpdateBody struct {
	FirstName string `json:"firstName" validate:"omitempty,min=1"`
	LastName  string `json:"lastName" validate:"omitempty,min=1"`
	Email     string `json:"email" validate:"omitempty,email"`
}

func (c *Controller) Update(ctx *fiber.Ctx) (err error) {
	var body UpdateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if body.FirstName == "" && body.LastName == "" && body.Email == "" {
		return fiber.ErrUnprocessableEntity
	}

	user := auth.GetSessionUser(ctx)

	if err = c.updateUser(ctx.Context(), user, &body); err != nil {
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

func (c *Controller) updateUser(ctx context.Context, user *model.User, body *UpdateBody) error {
	updated, err := c.workosClient.UserManagement.UpdateUser(ctx, usermanagement.UpdateUserOpts{
		User:      user.WorkosID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	})

	if err != nil {
		return err
	}

	return c.users.Update(ctx,
		model.User{
			Email:     updated.Email,
			FirstName: updated.FirstName,
			LastName:  updated.LastName,
		},
		repository.IDEq(user.ID),
	)
}
