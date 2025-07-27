package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/mfa"

	"polimane/backend/api/base"
)

type createAuthFactorBody struct {
	ChallengeID string `json:"challengeId" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

var (
	ErrInvalidAuthFactor = base.NewReasonedError(fiber.StatusBadRequest, "InvalidAuthFactor")
)

func (c *Controller) apiAuthFactorCreate(ctx *fiber.Ctx) (err error) {
	var body createAuthFactorBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	response, err := c.workosClient.MFA.VerifyChallenge(ctx.Context(), mfa.VerifyChallengeOpts{
		Code:        body.Code,
		ChallengeID: body.ChallengeID,
	})

	if err != nil {
		return err
	}

	if !response.Valid {
		return ErrInvalidAuthFactor
	}

	factor, err := c.workosClient.MFA.GetFactor(ctx.Context(), mfa.GetFactorOpts{
		FactorID: response.Challenge.FactorID,
	})

	return ctx.JSON(factor)
}
