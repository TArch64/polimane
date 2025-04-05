package repositoryschemas

import (
	"context"

	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
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
			SK: model.NewID(model.SKSchema).Key(),
		},
		Content: options.Content,
	}

	err := awsdynamodb.Table().Put(schema).Run(ctx)
	return schema, err
}
