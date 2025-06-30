package schemas

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type updateBody struct {
	Name    string              `json:"name" validate:"omitempty,min=1"`
	Palette []string            `json:"palette" validate:"omitempty,len=9,dive,omitempty,iscolor"`
	Content model.SchemaContent `json:"content" validate:"omitempty,dive,required"`
}

func collectUpdates(body *updateBody) model.Updates {
	updates := model.NewUpdates()

	if len(body.Name) > 0 {
		updates = updates.Set("Name", body.Name)
	}

	if body.Content != nil {
		updates = updates.Set("Content", body.Content)
	}

	if len(body.Palette) == repositoryschemas.PaletteSize {
		updates = updates.Set("Palette", body.Palette)
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

	err = repositoryschemas.Update(ctx.Context(), user, model.NewID(model.PKSchemaPrefix, schemaId), updates)
	if errors.Is(err, dynamo.ErrNotFound) {
		return base.NotFoundErr
	}
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
