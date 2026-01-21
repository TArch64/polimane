package schemas

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
)

type UpdateBody struct {
	Name            string              `json:"name" validate:"omitempty,min=1,max=255"`
	BackgroundColor string              `json:"backgroundColor" validate:"omitempty,iscolor"`
	Palette         model.SchemaPalette `json:"palette" validate:"omitempty,dive,omitempty,iscolor"`
	Size            *model.SchemaSize   `json:"size" validate:"omitempty"`
	Beads           model.SchemaBeads   `json:"beads" validate:"omitempty"`
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, ParamSchemaID)
	if err != nil {
		return err
	}

	var body UpdateBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	user := auth.GetSessionUser(ctx)

	userSchema, err := c.userSchemas.Get(reqCtx,
		repository.UserIDEq(user.ID),
		repository.SchemaIDEq(schemaID),
		repository.AccessGTE(model.AccessWrite),
	)
	if err != nil {
		return err
	}

	updates := collectUpdates(&body)
	if updates == nil {
		return base.NewReasonedError(fiber.StatusBadRequest, "EmptyUpdatesInput")
	}

	var beadsCounter *uint16
	if beads := updates.Beads.Data(); beads != nil {
		beadsLen := uint16(len(beads))
		if c.subscriptionCounters.SchemaBeads.CanSet(user, beadsLen) {
			beadsCounter = &beadsLen
		} else {
			return base.SchemasCreatedLimitReachedErr
		}
	}

	err = c.schemas.DB.
		WithContext(reqCtx).
		Transaction(func(tx *gorm.DB) error {
			err = c.schemas.UpdateTx(reqCtx, tx, *updates,
				repository.IDEq(schemaID),
			)
			if err != nil {
				return err
			}

			if beadsCounter != nil {
				return c.subscriptionCounters.SchemaBeads.SetTx(reqCtx, tx, userSchema, *beadsCounter)
			}

			return nil
		})

	if err != nil {
		return err
	}

	needImmediateScreenshotUpdate := body.BackgroundColor != ""
	if err = c.updateScreenshot(reqCtx, schemaID, needImmediateScreenshotUpdate); err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func collectUpdates(body *UpdateBody) *model.Schema {
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
			DeduplicationID: schemaID.String(),

			Body: events.SchemaScreenshotBody{
				SchemaID: schemaID,
			},
		})
	}

	schema, err := c.schemas.Get(ctx, repository.IDEq(schemaID))
	if err != nil {
		return err
	}

	return c.schemaScreenshot.Screenshot(ctx, &schemascreenshot.ScreenshotOptions{
		Schema: schema,
	})
}
