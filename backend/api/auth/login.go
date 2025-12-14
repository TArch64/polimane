package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {
	url, err := c.workosClient.UserManagement.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID:    c.env.WorkOS.ClientID,
		RedirectURI: c.env.ApiURL.JoinPath("api/auth/login/complete").String(),
		Provider:    "authkit",
	})

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"url": url.String(),
	})
}
