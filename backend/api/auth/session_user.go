package auth

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/model"
)

type sessionUserKeyType struct{}

var sessionUserKey sessionUserKeyType

func setSessionUser(ctx *fiber.Ctx, user *model.User) {
	ctx.Locals(sessionUserKey, user)
}

func GetSessionUser(ctx *fiber.Ctx) *model.User {
	return ctx.Locals(sessionUserKey).(*model.User)
}
