package schemas

import (
	"github.com/gofiber/fiber/v2"

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

func collectUpdates(body *updateBody) *model.Schema {
	changed := false
	updates := &model.Schema{}

	if len(body.Name) > 0 {
		changed = true
		updates.Name = body.Name
	}

	if body.Content != nil {
		changed = true
		updates.Content = body.Content
	}

	if body.Palette != nil {
		changed = true
		updates.Palette = body.Palette
	}

	if changed {
		return updates
	}

	return nil
}

func (c *Controller) apiUpdate(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, "schemaId")
	if err != nil {
		return err
	}

	var body updateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	updates := collectUpdates(&body)
	if updates == nil {
		return base.NewReasonedError(fiber.StatusBadRequest, "EmptyUpdatesInput")
	}

	err = c.schemas.Update(&repositoryschemas.UpdateOptions{
		Ctx:      ctx.Context(),
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
		Updates:  updates,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
