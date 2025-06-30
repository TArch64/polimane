package repositoryschemas

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

type ByIDOptions struct {
	Ctx        context.Context
	User       *model.User
	ID         model.ID
	Attributes []string
}

func ByID(options *ByIDOptions) (*model.Schema, error) {
	var err error
	if err = options.User.CheckSchemaAccess(options.ID); err != nil {
		return nil, err
	}

	var schema model.Schema

	query := awsdynamodb.Table().
		Get("PK", options.ID).
		Range("SK", dynamo.Equal, model.SKSchema)

	if len(options.Attributes) > 0 {
		options.Attributes = append([]string{"PK", "SK"}, options.Attributes...)
		query = query.Project(options.Attributes...)
	}

	err = query.One(options.Ctx, &schema)
	return &schema, err
}
