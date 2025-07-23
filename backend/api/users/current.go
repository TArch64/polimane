package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model/modelbase"
)

type currentUser struct {
	ID                modelbase.ID `json:"id"`
	FirstName         string       `json:"firstName"`
	LastName          string       `json:"lastName"`
	Username          string       `json:"username"`
	Email             string       `json:"email"`
	EmailVerified     bool         `json:"isEmailVerified"`
	ProfilePictureURL string       `json:"profilePictureUrl"`
}

func (c *Controller) apiCurrent(ctx *fiber.Ctx) error {
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
