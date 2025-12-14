package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

func (c *Controller) Logout(ctx *fiber.Ctx) error {
	err := Logout(ctx, &LogoutOptions{
		Env:          c.env,
		Signals:      c.signals,
		WorkosClient: c.workosClient,
	})
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

type LogoutOptions struct {
	Env          *env.Environment
	Signals      *signal.Container
	WorkosClient *workos.Client
}

func Logout(ctx *fiber.Ctx, options *LogoutOptions) error {
	session := GetSession(ctx)

	err := options.WorkosClient.UserManagement.RevokeSession(ctx.Context(), usermanagement.RevokeSessionOpts{
		SessionID: session.ID,
	})

	if err != nil {
		return err
	}

	removeCookies(ctx, options.Env)
	options.Signals.InvalidateAuthCache.Emit(ctx.Context(), session.ID)
	return nil
}
