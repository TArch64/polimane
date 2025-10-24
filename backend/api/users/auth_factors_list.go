package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
)

type authFactorListItem struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func (c *Controller) apiListAuthFactors(ctx *fiber.Ctx) error {
	user := auth.GetSessionUser(ctx)

	factors, err := c.workosClient.UserManagement.ListAuthFactors(ctx.Context(), usermanagement.ListAuthFactorsOpts{
		User: user.WorkosID,
	})

	if err != nil {
		return err
	}

	response := make([]authFactorListItem, len(factors.Data))
	for i, factor := range factors.Data {
		response[i] = authFactorListItem{
			ID:        factor.ID,
			CreatedAt: factor.CreatedAt,
		}
	}

	return ctx.JSON(response)
}
