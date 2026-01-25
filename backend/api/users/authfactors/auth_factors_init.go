package authfactors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
)

type NewAuthFactorResponse struct {
	ChallengeID string `json:"challengeId"`
	QRCode      string `json:"qrCode"`
	Secret      string `json:"secret"`
	URI         string `json:"uri"`
}

func (c *Controller) AuthFactorsInit(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)

	response, err := c.workos.UserManagement.EnrollAuthFactor(ctx.Context(), usermanagement.EnrollAuthFactorOpts{
		User:       session.WorkosUser.ID,
		Type:       mfa.TOTP,
		TOTPIssuer: "Polimane",
		TOTPUser:   session.WorkosUser.Email,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(NewAuthFactorResponse{
		ChallengeID: response.Challenge.ID,
		QRCode:      response.Factor.TOTP.QRCode,
		Secret:      response.Factor.TOTP.Secret,
		URI:         response.Factor.TOTP.URI,
	})
}
