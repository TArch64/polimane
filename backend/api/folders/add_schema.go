package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type addBody struct {
	SchemaIDs []model.ID `json:"schemaIds" validate:"dive,required"`
}

func (c *Controller) apiAddSchema(ctx *fiber.Ctx) (err error) {
	var body addBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	folderId, err := base.GetParamID(ctx, folderIDParam)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	err = c.folders.HasAccess(requestCtx, user.ID, folderId)
	if err != nil {
		return err
	}

	err = c.userSchemas.FilterByAccess(requestCtx, user, &body.SchemaIDs, model.AccessRead)
	if err != nil {
		return err
	}

	err = c.userSchemas.Update(requestCtx,
		model.UserSchema{FolderID: &folderId},
		repository.SchemaIDsIn(body.SchemaIDs),
	)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
