package users

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type updateAccessBody struct {
	bulkOperationBody
	Access model.AccessLevel `validate:"required,gte=1,lte=3" json:"access"`
}

func (c *Controller) apiUpdateAccess(ctx *fiber.Ctx) error {
	userID, err := base.GetParamID(ctx, userIDParam)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == userID {
		return base.InvalidRequestErr
	}

	var body updateAccessBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(requestCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return err
	}

	err = c.updateUserAccess(requestCtx, userID, body.IDs, body.Access)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) updateUserAccess(
	ctx context.Context,
	userID model.ID,
	schemaIDs []model.ID,
	access model.AccessLevel,
) error {
	return c.userSchemas.UpsertMany(ctx, &repositoryuserschemas.UpsertManyOptions{
		UserID:    userID,
		SchemaIDs: schemaIDs,
		Updates:   &model.UserSchema{Access: access},
	})
}
