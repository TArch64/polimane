package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
)

type newAuthFactorResponse struct {
	ChallengeID string `json:"challengeId"`
	QRCode      string `json:"qrCode"`
	Secret      string `json:"secret"`
	URI         string `json:"uri"`
}

func (c *Controller) apiAuthFactorsInit(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)

	response, err := c.workosClient.UserManagement.EnrollAuthFactor(ctx.Context(), usermanagement.EnrollAuthFactorOpts{
		User:       session.WorkosUser.ID,
		Type:       mfa.TOTP,
		TOTPIssuer: "Polimane",
		TOTPUser:   session.WorkosUser.Email,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(newAuthFactorResponse{
		ChallengeID: response.Challenge.ID,
		QRCode:      response.Factor.TOTP.QRCode,
		Secret:      response.Factor.TOTP.Secret,
		URI:         response.Factor.TOTP.URI,
	})
}
