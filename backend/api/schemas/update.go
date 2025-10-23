package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
)

type updateBody struct {
	Name            string              `json:"name" validate:"omitempty,min=1"`
	BackgroundColor string              `json:"backgroundColor" validate:"omitempty,iscolor"`
	Palette         model.SchemaPalette `json:"palette" validate:"omitempty,dive,omitempty,iscolor"`
	Size            *model.SchemaSize   `json:"size" validate:"omitempty"`
	Beads           model.SchemaBeads   `json:"beads" validate:"omitempty"`
}

func collectUpdates(body *updateBody) *model.Schema {
	changed := false
	updates := &model.Schema{}

	if body.Name != "" {
		changed = true
		updates.Name = body.Name
	}

	if body.BackgroundColor != "" {
		changed = true
		updates.BackgroundColor = body.BackgroundColor
	}

	if body.Palette != nil {
		changed = true
		updates.Palette = datatypes.NewJSONType(body.Palette)
	}

	if body.Size != nil {
		changed = true
		updates.Size = datatypes.NewJSONType(body.Size)
	}

	if body.Beads != nil {
		changed = true
		updates.Beads = datatypes.NewJSONType(body.Beads)
	}

	if changed {
		return updates
	}

	return nil
}

func (c *Controller) updateScreenshot(ctx context.Context, schemaID model.ID, needImmediateUpdate bool) error {
	if !needImmediateUpdate {
		return c.sqs.Send(ctx, &awssqs.SendOptions{
			Queue:           events.QueueDebounced,
			Event:           events.EventSchemaScreenshot,
			DeduplicationId: schemaID.String(),

			Body: events.SchemaScreenshotBody{
				SchemaID: schemaID,
			},
		})
	}

	schema, err := c.schemas.GetByID(ctx, &repositoryschemas.ByIDOptions{
		SchemaID: schemaID,
	})

	if err != nil {
		return err
	}

	return c.schemaScreenshot.Screenshot(ctx, &schemascreenshot.ScreenshotOptions{
		Schema: schema,
	})
}

func (c *Controller) apiUpdate(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIdParam)
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

	err = c.schemas.Update(ctx.Context(), &repositoryschemas.UpdateOptions{
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaID,
		Updates:  updates,
	})

	if err != nil {
		return err
	}

	needImmediateScreenshotUpdate := body.BackgroundColor != ""
	if err = c.updateScreenshot(ctx.Context(), schemaID, needImmediateScreenshotUpdate); err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
