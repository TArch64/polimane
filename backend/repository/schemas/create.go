package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	repositoryusers "polimane/backend/repository/users"
	awsdynamodb "polimane/backend/services/dynamodb"
	"polimane/backend/signal"
)

type CreateOptions struct {
	Ctx     context.Context
	User    *model.User
	Name    string
	Palette []string
	Content model.SchemaContent
}

const PaletteSize = 9

func Create(options *CreateOptions) (*model.Schema, error) {
	if len(options.Palette) == 0 {
		options.Palette = make([]string, PaletteSize)
	}

	if options.Content == nil {
		options.Content = make(model.SchemaContent, 0)
	}

	schema := &model.Schema{
		Base: &model.Base{
			PK: model.RandomID(model.PKSchemaPrefix),
			SK: model.SKSchema,
		},
		UserIDs: []model.PrimaryKey{options.User.PrimaryKey()},
		Name:    options.Name,
		Palette: options.Palette,
		Content: options.Content,
	}

	userUpdates := model.NewUpdates().Add("SchemaIDs", schema.PrimaryKey().String())

	tx := awsdynamodb.WriteTX().
		Put(awsdynamodb.Table().Put(schema)).
		Update(repositoryusers.UpdateTx(options.User.PrimaryKey(), userUpdates))

	if err := tx.Run(options.Ctx); err != nil {
		return nil, err
	}

	options.User.AddSchemaID(schema.PrimaryKey())
	signal.InvalidateAuthCache.Emit(options.Ctx, options.User.PK)
	return schema, nil
}
