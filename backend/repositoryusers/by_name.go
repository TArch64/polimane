package repositoryusers

import (
	"context"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func ByName(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	err := awsdynamodb.Table().
		Get("SK", model.NewKey(model.SKUser, username)).
		Index(model.IndexUserName).
		One(ctx, &user)

	return &user, err
}
