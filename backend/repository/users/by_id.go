package repositoryusers

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func ByID(ctx context.Context, id model.ID) (*model.User, error) {
	var user model.User

	err := awsdynamodb.Table().
		Get("PK", id).
		Range("SK", dynamo.BeginsWith, model.SKUserPrefix).
		One(ctx, &user)

	return &user, err
}
