package authfactors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
)

type AuthFactorListItem struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func (c *Controller) ListAuthFactors(ctx *fiber.Ctx) error {
	user := auth.GetSessionUser(ctx)

	factors, err := c.workos.UserManagement.ListAuthFactors(ctx.Context(), usermanagement.ListAuthFactorsOpts{
		User: user.WorkosID,
	})

	if err != nil {
		return err
	}

	response := make([]AuthFactorListItem, len(factors.Data))
	for i, factor := range factors.Data {
		response[i] = AuthFactorListItem{
			ID:        factor.ID,
			CreatedAt: factor.CreatedAt,
		}
	}

	return ctx.JSON(response)
}
