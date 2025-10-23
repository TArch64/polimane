package schemas

import (
	"context"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"polimane/backend/model"
)

type CreateOptions struct {
	User            *model.User
	Name            string
	BackgroundColor string
	Palette         model.SchemaPalette
	Size            *model.SchemaSize
	Beads           model.SchemaBeads
}

func (i *Client) Create(ctx context.Context, options *CreateOptions) (schema *model.Schema, err error) {
	if options.Palette == nil {
		options.Palette = make(model.SchemaPalette, model.SchemaPaletteSize)
	}

	if options.Size == nil {
		options.Size = &model.SchemaSize{
			Left:   50,
			Right:  49,
			Top:    15,
			Bottom: 14,
		}
	}

	if options.Beads == nil {
		options.Beads = make(model.SchemaBeads)
	}

	err = i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		schema = &model.Schema{
			Name:            options.Name,
			BackgroundColor: options.BackgroundColor,
			Palette:         datatypes.NewJSONType(options.Palette),
			Size:            datatypes.NewJSONType(options.Size),
			Beads:           datatypes.NewJSONType(options.Beads),
		}

		if err = tx.Create(schema).Error; err != nil {
			return err
		}

		return i.userSchemas.CreateTx(tx, options.User.ID, schema.ID, model.AccessAdmin)
	})

	if err != nil {
		return nil, err
	}

	i.signals.InvalidateUserCache.Emit(ctx, options.User.ID)
	return schema, nil
}
