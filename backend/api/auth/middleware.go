package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/base"
	repositoryusers "polimane/backend/repository/users"
)

var unauthorizedErr = base.NewReasonedError(fiber.StatusUnauthorized, "Unauthorized")

func Middleware(ctx *fiber.Ctx) error {
	token := ctx.Cookies(cookieName)

	if len(token) == 0 {
		return unauthorizedErr
	}

	claims, err := parseCookieToken(token)
	if err != nil {
		return err
	}

	user, err := repositoryusers.ByPK(ctx.Context(), claims.UserID)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return unauthorizedErr
		}
		return err
	}

	setSessionUser(ctx, user)
	return ctx.Next()
}
