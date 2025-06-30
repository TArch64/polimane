package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	repositoryusers "polimane/backend/repository/users"
	awsdynamodb "polimane/backend/services/dynamodb"
	"polimane/backend/signal"
)

func Delete(ctx context.Context, user *model.User, id model.ID) (err error) {
	if err = user.CheckSchemaAccess(id); err != nil {
		return err
	}

	schema, err := ByID(&ByIDOptions{
		Ctx:        ctx,
		ID:         id,
		User:       user,
		Attributes: []string{"UserIDs"},
	})

	if err != nil {
		return err
	}

	del := awsdynamodb.Table().
		Delete("PK", id).
		Range("SK", model.SKSchema)

	userUpdates := model.NewUpdates().
		Delete("SchemaIDs", schema.PrimaryKey().String())

	tx := awsdynamodb.WriteTX().
		Delete(del)

	for _, userID := range schema.UserIDs {
		tx.Update(repositoryusers.UpdateTx(userID, userUpdates))
	}

	if err = tx.Run(ctx); err != nil {
		return model.ConditionErrToNotFound(err)
	}

	user.DeleteSchemaID(id)
	signal.InvalidateAuthCache.Emit(ctx, user.PK)
	return nil
}
