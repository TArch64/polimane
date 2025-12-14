package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type AddBody struct {
	SchemaIDs []model.ID `json:"schemaIds" validate:"dive,required"`
}

func (c *Controller) AddSchema(ctx *fiber.Ctx) (err error) {
	var body AddBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	folderID, err := base.GetParamID(ctx, ParamFolderID)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	reqCtx := ctx.Context()
	err = c.folders.HasAccess(reqCtx, user.ID, folderID)
	if err != nil {
		return err
	}

	err = c.userSchemas.FilterByAccess(reqCtx, user, &body.SchemaIDs, model.AccessRead)
	if err != nil {
		return err
	}

	err = c.userSchemas.Update(reqCtx,
		model.UserSchema{FolderID: &folderID},
		repository.SchemaIDsIn(body.SchemaIDs),
	)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
