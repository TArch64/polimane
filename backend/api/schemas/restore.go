package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func (c *Controller) Restore(ctx *fiber.Ctx) (err error) {
	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	if err = c.schemas.Restore(reqCtx, body.IDs); err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	err = c.subscriptionCounters.SchemasCreated.Add(reqCtx,
		uint16(len(body.IDs)),
		user.ID,
	)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
