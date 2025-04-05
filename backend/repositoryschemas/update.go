package repositoryschemas

import (
	"context"

	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
)

func Update(ctx context.Context, user *model.User, id string, patch awsdynamodb.UpdateMap) error {
	query := awsdynamodb.Table().
		Update("PK", user.ID).
		Range("SK", model.NewID(model.SKSchema, id))

	for key, value := range patch {
		query = query.Set(key, value)
	}

	return query.Run(ctx)
}
