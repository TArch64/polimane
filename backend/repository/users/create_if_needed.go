package repositoryusers

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/services/db"
)

func CreateIfNeeded(ctx context.Context, workosID string) (*model.User, error) {
	user := &model.User{WorkosID: workosID}

	err := db.Instance.
		WithContext(ctx).
		Where(*user).
		Attrs(*user).
		FirstOrCreate(user).
		Error

	return user, err
}
