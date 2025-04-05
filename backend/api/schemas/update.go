package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
	"polimane/backend/repositoryschemas"
)

type updateBody struct {
	Name    string              `json:"name"`
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
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
