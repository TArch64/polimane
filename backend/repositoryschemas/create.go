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
			SK: model.RandomID(model.SKSchema).Key(),
		},
		Name:    options.Name,
		Content: options.Content,
	}

	err := awsdynamodb.Table().
		Put(schema).
		If("attribute_not_exists(PR)").
		Run(ctx)

	return schema, err
}
