package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
)

type UserResponse struct {
	ID            model.ID `json:"id"`
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"isEmailVerified"`
}

func (c *Controller) Current(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)

	return ctx.JSON(UserResponse{
		ID:            session.User.ID,
		FirstName:     session.User.FirstName,
		LastName:      session.User.LastName,
		Email:         session.User.Email,
		EmailVerified: session.WorkosUser.EmailVerified,
	})
}
