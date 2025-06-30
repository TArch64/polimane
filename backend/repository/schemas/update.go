package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	awsdynamodb "polimane/backend/services/dynamodb"
)

func Update(ctx context.Context, user *model.User, id model.ID, updates model.Updates) (err error) {
	if err = user.CheckSchemaAccess(id); err != nil {
		return err
	}

	schema, err := ByID(&ByIDOptions{
		Ctx:        ctx,
		ID:         id,
		User:       user,
		Attributes: make([]string, 0),
	})

	if err != nil {
		return err
	}

	update := awsdynamodb.Table().
		Update("PK", schema.PK).
		Range("SK", schema.SK)

	return updates.Apply(update).Run(ctx)
}
