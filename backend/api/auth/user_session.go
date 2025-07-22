package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
)

type UserSession struct {
	User       *model.User
	WorkosUser *usermanagement.User
	SessionID  string
}

func (u *UserSession) GetSID() string {
	return u.WorkosUser.Metadata["SessionID"]
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

func GetSessionID(ctx *fiber.Ctx) string {
	return GetSession(ctx).SessionID
}
