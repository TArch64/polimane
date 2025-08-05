package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
)

type currentUser struct {
	ID                model.ID `json:"id"`
	FirstName         string   `json:"firstName"`
	LastName          string   `json:"lastName"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"isEmailVerified"`
	ProfilePictureURL string   `json:"profilePictureUrl"`
}

func (c *Controller) apiGet(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)

	return ctx.JSON(currentUser{
		ID:                session.User.ID,
		FirstName:         session.WorkosUser.FirstName,
		LastName:          session.WorkosUser.LastName,
		Email:             session.WorkosUser.Email,
		EmailVerified:     session.WorkosUser.EmailVerified,
		ProfilePictureURL: session.WorkosUser.ProfilePictureURL,
	})
}
