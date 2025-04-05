package repositoryschemas

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
)

func ByUser(ctx context.Context, user *model.User) ([]*model.Schema, error) {
	var schemas []*model.Schema

	err := awsdynamodb.Table().
		Get("PK", user.ID).
		Range("SK", dynamo.BeginsWith, model.SKSchema+"#").
		All(ctx, &schemas)

	return schemas, err
}
