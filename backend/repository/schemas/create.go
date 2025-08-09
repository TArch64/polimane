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
	Content *model.SchemaContent
}

func (i *Impl) Create(options *CreateOptions) (schema *model.Schema, err error) {
	if options.Palette == nil {
		options.Palette = make(model.SchemaPalette, model.SchemaPaletteSize)
	}

	if options.Content == nil {
		options.Content = i.createDefaultContent()
	}

	err = i.db.WithContext(options.Ctx).Transaction(func(tx *gorm.DB) error {
		schema = &model.Schema{
			Name:    options.Name,
			Palette: datatypes.NewJSONType(options.Palette),
			Content: datatypes.NewJSONType(options.Content),
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

func (i *Impl) createDefaultContent() *model.SchemaContent {
	return &model.SchemaContent{
		Size: &model.SchemaContentSize{
			Left:   25,
			Top:    25,
			Right:  24,
			Bottom: 24,
		},
		Beads: make(map[string]string),
	}
}
