package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/repository"
)

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
	folderID, err := base.GetParamID(ctx, folderIDParam)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	err = c.folders.Delete(ctx.Context(),
		repository.IDEq(folderID),
		repository.UserIDEq(currentUser.ID),
	)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
