package repositoryschemas

import (
	"context"

	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
)

func Delete(ctx context.Context, user *model.User, id string) error {
	return awsdynamodb.Table().
		Delete("PK", user.ID).
		Range("SK", model.NewID(model.SKSchema, id)).
		Run(ctx)
}
