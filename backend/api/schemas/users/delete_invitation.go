package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/repository"
)

func (c *Controller) DeleteInvitation(ctx *fiber.Ctx) (err error) {
	var body InvitationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.FilterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	err = c.schemaInvitations.Delete(ctx.Context(),
		repository.EmailEq(body.Email),
		repository.SchemaIDsIn(body.IDs),
	)

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
