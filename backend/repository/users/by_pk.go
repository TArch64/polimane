package repositoryusers

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func ByPK(ctx context.Context, pk model.ID) (*model.User, error) {
	var user model.User

	err := awsdynamodb.Table().
		Get("PK", pk).
		Range("SK", dynamo.BeginsWith, model.SKUser).
		One(ctx, &user)

	return &user, err
}
