package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/mfa"

	"polimane/backend/api/base"
)

func (c *Controller) AuthFactorDelete(ctx *fiber.Ctx) error {
	factorID, err := base.GetRequiredParam(ctx, ParamFactorID)
	if err != nil {
		return err
	}

	err = c.workosClient.MFA.DeleteFactor(ctx.Context(), mfa.DeleteFactorOpts{
		FactorID: factorID,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
