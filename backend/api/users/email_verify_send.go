package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func (c *Controller) apiEmailVerifyRetry(ctx *fiber.Ctx) (err error) {
	user := auth.GetSessionUser(ctx)

	if err = c.sendEmailVerification(ctx.Context(), user.WorkosID); err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) sendEmailVerification(ctx context.Context, userID string) error {
	_, err := c.workosClient.UserManagement().SendVerificationEmail(ctx, usermanagement.SendVerificationEmailOpts{
		User: userID,
	})

	return err
}
