package repositoryschemas

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func Update(ctx context.Context, user *model.User, id string, patch awsdynamodb.UpdateMap) error {
	query := awsdynamodb.Table().
		Update("PK", user.ID).
		Range("SK", model.NewID(model.SKSchema, id)).
		If("attribute_exists(PK)")

	for key, value := range patch {
		query = query.Set(key, value)
	}

	err := query.Run(ctx)

	var checkFailedErr *types.ConditionalCheckFailedException
	if errors.As(err, &checkFailedErr) {
		return dynamo.ErrNotFound
	}

	return err
}
