package repositoryschemas

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func ById(ctx context.Context, user *model.User, id string) (*model.Schema, error) {
	var schema model.Schema

	err := awsdynamodb.Table().
		Get("PK", user.ID).
		Range("SK", dynamo.Equal, model.NewID(model.SKSchema, id)).
		One(ctx, &schema)

	return &schema, err
}
