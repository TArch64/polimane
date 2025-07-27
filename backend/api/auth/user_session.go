package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
)

type UserSession struct {
	ID         string
	User       *model.User
	WorkosUser *usermanagement.User
}

var sessionKey UserSession

func setSession(ctx *fiber.Ctx, session *UserSession) {
	ctx.Locals(sessionKey, session)
}

func GetSession(ctx *fiber.Ctx) *UserSession {
	return ctx.Locals(sessionKey).(*UserSession)
}

func GetSessionUser(ctx *fiber.Ctx) *model.User {
	return GetSession(ctx).User
}

func GetSessionWorkosUser(ctx *fiber.Ctx) *usermanagement.User {
	return GetSession(ctx).WorkosUser
}

func GetSessionID(ctx *fiber.Ctx) string {
	return GetSession(ctx).ID
}
