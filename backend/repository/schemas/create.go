package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

type CreateOptions struct {
	User    *model.User
	Name    string
	Palette []string
	Content model.SchemaContent
}

const PaletteSize = 9

func Create(ctx context.Context, options *CreateOptions) (*model.Schema, error) {
	if len(options.Palette) == 0 {
		options.Palette = make([]string, PaletteSize)
	}

	if options.Content == nil {
		options.Content = make(model.SchemaContent, 0)
	}

	schema := &model.Schema{
		Base: &model.Base{
			ID: options.User.ID,
			SK: model.RandomID(model.SKSchema).Key(),
		},
		Name:    options.Name,
		Palette: options.Palette,
		Content: options.Content,
	}

	err := awsdynamodb.Table().Put(schema).Run(ctx)
	return schema, err
}
