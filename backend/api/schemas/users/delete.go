package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
	userID, err := base.GetParamID(ctx, userIDParam)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == userID {
		return base.InvalidRequestErr
	}

	var body bulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(requestCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.userSchemas.Delete(
		requestCtx,
		repository.UserIDEq(userID),
		repository.SchemaIDsIn(body.IDs),
	)

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
