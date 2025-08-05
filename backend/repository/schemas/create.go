package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type CreateOptions struct {
	Ctx     context.Context
	User    *model.User
	Name    string
	Palette model.SchemaPalette
	Content model.SchemaContent
}

func (c *Impl) Create(options *CreateOptions) (schema *model.Schema, err error) {
	if options.Palette == nil {
		options.Palette = make(model.SchemaPalette, model.SchemaPaletteSize)
	}

	if options.Content == nil {
		options.Content = make(model.SchemaContent, 0)
	}

	err = c.db.WithContext(options.Ctx).Transaction(func(tx *gorm.DB) error {
		schema = &model.Schema{
			Name:    options.Name,
			Palette: options.Palette,
			Content: options.Content,
		}

		if err = tx.Create(schema).Error; err != nil {
			return err
		}

		return c.userSchemas.CreateTx(tx, options.User.ID, schema.ID)
	})

	if err != nil {
		return nil, err
	}

	c.signals.InvalidateUserCache.Emit(options.Ctx, options.User.ID)
	return schema, nil
}
