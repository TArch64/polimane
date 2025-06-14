package schemas

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	awsdynamodb "polimane/backend/services/dynamodb"
)

type updateBody struct {
	Name    string              `json:"name" validate:"omitempty,min=1"`
	Palette []string            `json:"palette" validate:"omitempty,len=9"`
	Content model.SchemaContent `json:"content"`
}

func collectUpdates(body *updateBody) awsdynamodb.UpdateMap {
	updates := awsdynamodb.UpdateMap{}

	if len(body.Name) > 0 {
		updates["Name"] = body.Name
	}

	if body.Content != nil {
		updates["Content"] = body.Content
	}

	if len(body.Palette) != 0 {
		updates["Palette"] = body.Palette
	}

	return updates
}

func apiUpdate(ctx *fiber.Ctx) error {
	schemaId, err := base.GetRequiredParam(ctx, "schemaId")
	if err != nil {
		return err
	}

	var body updateBody
	err = base.ParseBody(ctx, &body)
	if err != nil {
		return err
	}

	updates := collectUpdates(&body)
	if len(updates) == 0 {
		return base.NewReasonedError(fiber.StatusBadRequest, "EmptyUpdatesInput")
	}

	user := auth.GetSessionUser(ctx)

	err = repositoryschemas.Update(ctx.Context(), user, schemaId, updates)
	if errors.Is(err, dynamo.ErrNotFound) {
		return base.NotFoundErr
	}
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
