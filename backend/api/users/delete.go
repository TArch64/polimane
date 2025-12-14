package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/repository"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	currentUser := auth.GetSessionUser(ctx)

	err := c.users.Delete(ctx.Context(),
		repository.IDEq(currentUser.ID),
	)
	if err != nil {
		return err
	}

	err = auth.Logout(ctx, &auth.LogoutOptions{
		Env:          c.env,
		Signals:      c.signals,
		WorkosClient: c.workosClient,
	})
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
