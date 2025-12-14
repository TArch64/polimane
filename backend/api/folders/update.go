package folders

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	dberror "polimane/backend/services/db/error"
)

type updateBody struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

func (c *Controller) apiUpdate(ctx *fiber.Ctx) error {
	folderID, err := base.GetParamID(ctx, folderIDParam)
	if err != nil {
		return err
	}

	var body updateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	err = c.folders.Update(ctx.Context(),
		model.Folder{Name: body.Name},
		repository.IDEq(folderID),
		repository.UserIDEq(user.ID),
	)

	if errors.Is(err, dberror.UniqueConstraintErr) {
		return NameAlreadyInUseErr
	}
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
