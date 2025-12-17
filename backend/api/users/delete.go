package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/repository"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	currentUser := auth.GetSessionUser(ctx)

	_, err := c.workos.UserManagement.UpdateUser(ctx.Context(), usermanagement.UpdateUserOpts{
		User: currentUser.WorkosID,

		Metadata: map[string]string{
			"ScheduledForDeletion": "true",
		},
	})
	if err != nil {
		return err
	}

	err = c.users.Delete(ctx.Context(),
		repository.IDEq(currentUser.ID),
	)
	if err != nil {
		return err
	}

	err = auth.Logout(ctx, &auth.LogoutOptions{
		Env:     c.env,
		Signals: c.signals,
		Workos:  c.workos,
	})
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
