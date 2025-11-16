package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type folderDetails struct {
	ID   model.ID `json:"id"`
	Name string   `json:"name"`
}

func (c *Controller) apiById(ctx *fiber.Ctx) error {
	folderId, err := base.GetParamID(ctx, folderIDParam)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	var folder folderDetails

	err = c.folders.GetOut(ctx.Context(), &folder,
		repository.Select("id", "name"),
		repository.IDEq(folderId),
		repository.UserIDEq(user.ID),
	)

	if err != nil {
		return err
	}

	return ctx.JSON(folder)
}
