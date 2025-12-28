package folders

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Controller) RemoveSchemas(ctx *fiber.Ctx) (err error) {
	var body BulkSchemasBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if _, err = c.getFolderID(ctx); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.SchemaIDs); err != nil {
		return err
	}

	err = c.userSchemas.Update(ctx.Context(),
		model.UserSchema{FolderID: model.NilID},
		repository.UserIDEq(auth.GetSessionUser(ctx).ID),
		repository.SchemaIDsIn(body.SchemaIDs),
	)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
