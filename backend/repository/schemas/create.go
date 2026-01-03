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
	Layout          model.SchemaLayout
	Palette         model.SchemaPalette
	Size            *model.SchemaSize
	Beads           model.SchemaBeads
	FolderID        *model.ID
}

func (c *Client) Create(ctx context.Context, options *CreateOptions) (schema *model.Schema, err error) {
	if options.Palette == nil {
		options.Palette = model.DefaultPalette
	}

	if options.Size == nil {
		size := model.DefaultSchemaSize
		options.Size = &size
	}

	if options.Beads == nil {
		options.Beads = make(model.SchemaBeads)
	}

	err = c.DB.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			schema = &model.Schema{
				Name:            options.Name,
				BackgroundColor: options.BackgroundColor,
				Layout:          options.Layout,
				Palette:         datatypes.NewJSONType(options.Palette),
				Size:            datatypes.NewJSONType(options.Size),
				Beads:           datatypes.NewJSONType(options.Beads),
			}

			if err = c.InsertTx(ctx, tx, schema); err != nil {
				return err
			}

			return c.userSchemas.InsertTx(ctx, tx, &model.UserSchema{
				UserID:   options.User.ID,
				SchemaID: schema.ID,
				FolderID: options.FolderID,
				Access:   model.AccessAdmin,
			})
		})

	if err != nil {
		return nil, err
	}

	return schema, nil
}
