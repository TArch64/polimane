package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
)

func (c *Controller) DeletePermanently(ctx *fiber.Ctx) (err error) {
	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	if err = c.schemaDelete.Delete(ctx.Context(), body.IDs); err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
