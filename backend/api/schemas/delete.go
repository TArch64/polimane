package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Controller) Delete(ctx *fiber.Ctx) (err error) {
	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	reqCtx := ctx.Context()
	if err = c.schemas.DeleteSoft(reqCtx, user, body.IDs); err != nil {
		return err
	}

	affectedUserIDs, err := c.getAffectedUsersOnDelete(reqCtx, body.IDs)
	if err != nil {
		return err
	}

	err = c.subscriptionCounters.SchemasCreated.Remove(reqCtx,
		uint16(len(body.IDs)),
		affectedUserIDs...,
	)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) getAffectedUsersOnDelete(ctx context.Context, schemaIDs []model.ID) (out []model.ID, err error) {
	err = c.userSchemas.ListOut(ctx, &out,
		repository.IncludeSoftDeleted,
		repository.Select("DISTINCT ON (user_id) user_id"),
		repository.SchemaIDsIn(schemaIDs),
	)
	return
}
