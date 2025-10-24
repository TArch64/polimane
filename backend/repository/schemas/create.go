package schemas

import (
	"context"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type CreateOptions struct {
	User            *model.User
	Name            string
	BackgroundColor string
	Palette         model.SchemaPalette
	Size            *model.SchemaSize
	Beads           model.SchemaBeads
}

func (c *Client) Create(ctx context.Context, options *CreateOptions) (schema *model.Schema, err error) {
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

	err = c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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

		_, err = c.userSchemas.CreateTx(ctx, tx, &repositoryuserschemas.CreateOptions{
			UserID:   options.User.ID,
			SchemaID: schema.ID,
			Access:   model.AccessAdmin,
		})

		return err
	})

	if err != nil {
		return nil, err
	}

	c.signals.InvalidateUserCache.Emit(ctx, options.User.ID)
	return schema, nil
}
