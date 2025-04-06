package repositoryschemas

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
)

func Delete(ctx context.Context, user *model.User, id string) error {
	err := awsdynamodb.Table().
		Delete("PK", user.ID).
		Range("SK", model.NewID(model.SKSchema, id)).
		If("attribute_exists(PK)").
		Run(ctx)

	var checkFailedErr *types.ConditionalCheckFailedException
	if errors.As(err, &checkFailedErr) {
		return dynamo.ErrNotFound
	}

	return err
}
