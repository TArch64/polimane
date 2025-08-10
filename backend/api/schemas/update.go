package schemas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type updateBody struct {
	Name    string              `json:"name" validate:"omitempty,min=1"`
	Palette model.SchemaPalette `json:"palette" validate:"omitempty,len=9,dive,omitempty,iscolor"`
	Size    *model.SchemaSize   `json:"size" validate:"omitempty"`
	Beads   model.SchemaBeads   `json:"beads" validate:"omitempty,dive,required"`
}

func collectUpdates(body *updateBody) *model.Schema {
	changed := false
	updates := &model.Schema{}

	if len(body.Name) > 0 {
		changed = true
		updates.Name = body.Name
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

func (c *Controller) apiUpdate(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
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
