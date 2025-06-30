package repositoryschemas

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

type ByUserOptions struct {
	Ctx        context.Context
	User       *model.User
	Attributes []string
}

func ByUser(options *ByUserOptions) ([]*model.Schema, error) {
	if len(options.User.SchemaIDs) == 0 {
		return nil, nil
	}

	batchKeys := make([]dynamo.Keyed, len(options.User.SchemaIDs))
	for i, key := range options.User.SchemaIDs {
		batchKeys[i] = key.Keys()
	}

	query := awsdynamodb.Table().
		Batch("PK", "SK").
		Get(batchKeys...)

	if len(options.Attributes) > 0 {
		options.Attributes = append([]string{"PK", "SK"}, options.Attributes...)
		query = query.Project(options.Attributes...)
	}

	var schemas []*model.Schema
	err := query.All(options.Ctx, &schemas)
	return schemas, err
}
