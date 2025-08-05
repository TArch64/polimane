package users

import (
	"context"

	"polimane/backend/model"
)

func (c *Impl) CreateIfNeeded(ctx context.Context, workosID string) (*model.User, error) {
	user := &model.User{WorkosID: workosID}

	err := c.db.
		WithContext(ctx).
		Where(*user).
		Attrs(*user).
		FirstOrCreate(user).
		Error

	return user, err
}
