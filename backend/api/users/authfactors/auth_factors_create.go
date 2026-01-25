package authfactors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/mfa"

	"polimane/backend/api/base"
)

type CreateAuthFactorBody struct {
	ChallengeID string `json:"challengeId" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

var (
	ErrInvalidAuthFactor = base.NewReasonedError(fiber.StatusBadRequest, "InvalidAuthFactor")
)

func (c *Controller) AuthFactorCreate(ctx *fiber.Ctx) (err error) {
	var body CreateAuthFactorBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	response, err := c.workos.MFA.VerifyChallenge(ctx.Context(), mfa.VerifyChallengeOpts{
		Code:        body.Code,
		ChallengeID: body.ChallengeID,
	})

	if err != nil {
		return err
	}

	if !response.Valid {
		return ErrInvalidAuthFactor
	}

	factor, err := c.workos.MFA.GetFactor(ctx.Context(), mfa.GetFactorOpts{
		FactorID: response.Challenge.FactorID,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(factor)
}
