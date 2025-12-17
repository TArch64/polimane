package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type FolderDetails struct {
	ID   model.ID `json:"id"`
	Name string   `json:"name"`
}

func (c *Controller) ByID(ctx *fiber.Ctx) error {
	folderID, err := base.GetParamID(ctx, ParamFolderID)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	var folder FolderDetails

	err = c.folders.GetOut(ctx.Context(), &folder,
		repository.Select("id", "name"),
		repository.IDEq(folderID),
		repository.UserIDEq(user.ID),
	)

	if err != nil {
		return err
	}

	return ctx.JSON(folder)
}
