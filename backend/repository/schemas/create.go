package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

type CreateOptions struct {
	User    *model.User
	Name    string
	Content model.SchemaContent
}

func Create(ctx context.Context, options *CreateOptions) (*model.Schema, error) {
	schema := &model.Schema{
		Base: &model.Base{
			ID: options.User.ID,
			SK: model.RandomID(model.SKSchema).Key(),
		},
		Name:    options.Name,
		Palette: make([]string, 10),
		Content: options.Content,
	}

	err := awsdynamodb.Table().Put(schema).Run(ctx)
	return schema, err
}
