package schemas

import (
	"context"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"polimane/backend/model"
)

type CreateOptions struct {
	Ctx     context.Context
	User    *model.User
	Name    string
	Palette model.SchemaPalette
	Size    *model.SchemaSize
	Beads   model.SchemaBeads
}

func (i *Impl) Create(options *CreateOptions) (schema *model.Schema, err error) {
	if options.Palette == nil {
		options.Palette = make(model.SchemaPalette, model.SchemaPaletteSize)
	}

	if options.Size == nil {
		options.Size = &model.SchemaSize{
			Left:   25,
			Top:    25,
			Right:  24,
			Bottom: 24,
		}
	}

	if options.Beads == nil {
		options.Beads = make(model.SchemaBeads)
	}

	err = i.db.WithContext(options.Ctx).Transaction(func(tx *gorm.DB) error {
		schema = &model.Schema{
			Name:    options.Name,
			Palette: datatypes.NewJSONType(options.Palette),
			Size:    datatypes.NewJSONType(options.Size),
			Beads:   datatypes.NewJSONType(options.Beads),
		}

		if err = tx.Create(schema).Error; err != nil {
			return err
		}

		return i.userSchemas.CreateTx(tx, options.User.ID, schema.ID)
	})

	if err != nil {
		return nil, err
	}

	i.signals.InvalidateUserCache.Emit(options.Ctx, options.User.ID)
	return schema, nil
}
