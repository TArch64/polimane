package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	userID, err := base.GetParamID(ctx, ParamUserID)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == userID {
		return base.InvalidRequestErr
	}

	var body base.BulkOperationBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(reqCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.userSchemas.Delete(
		reqCtx,
		repository.UserIDEq(userID),
		repository.SchemaIDsIn(body.IDs),
	)

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
