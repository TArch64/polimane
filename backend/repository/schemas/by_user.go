package repositoryschemas

import (
	"context"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func ByUser(ctx context.Context, user *model.User, attributes []string) ([]*model.Schema, error) {
	var schemas []*model.Schema

	query := awsdynamodb.Table().
		Get("PK", user.ID).
		Range("SK", dynamo.BeginsWith, model.SKSchema+"#")

	if attributes != nil && len(attributes) > 0 {
		attributes = append([]string{"PK", "SK"}, attributes...)
		query = query.Project(attributes...)
	}

	err := query.All(ctx, &schemas)
	return schemas, err
}
