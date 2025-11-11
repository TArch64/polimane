package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

type deleteInvitationBody struct {
	bulkOperationBody
	Email string `json:"email" validate:"required,email,max=255"`
}

func (c *Controller) apiDeleteInvitation(ctx *fiber.Ctx) error {
	var err error
	var body deleteInvitationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	currentUser := auth.GetSessionUser(ctx)
	err = c.userSchemas.FilterByAccess(requestCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}
	if len(body.IDs) == 0 {
		return fiber.ErrBadRequest
	}

	err = c.schemaInvitations.DeleteMany(requestCtx,
		repository.EmailEq(body.Email),
		repository.SchemaIDsIn(body.IDs),
	)

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
