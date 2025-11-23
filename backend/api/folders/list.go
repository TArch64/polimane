package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/repository"
)

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	currentUser := auth.GetSessionUser(ctx)

	folders, err := c.folders.List(ctx.Context(),
		repository.UserIDEq(currentUser.ID),
		repository.Order("created_at DESC"),
	)
	if err != nil {
		return err
	}

	return ctx.JSON(folders)
}
