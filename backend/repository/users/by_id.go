package users

import (
	"context"

	"polimane/backend/model"
)

func (i *Impl) ByID(ctx context.Context, id model.ID) (*model.User, error) {
	var user model.User
	err := i.db.WithContext(ctx).Take(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
