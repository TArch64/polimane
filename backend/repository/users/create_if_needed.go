package users

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
)

func (i *Impl) CreateIfNeeded(ctx context.Context, workosUser usermanagement.User) (*model.User, error) {
	user := model.User{
		WorkosID:  workosUser.ID,
		Email:     workosUser.Email,
		FirstName: workosUser.FirstName,
		LastName:  workosUser.LastName,
	}

	err := i.db.
		WithContext(ctx).
		Where(user).
		Attrs(user).
		FirstOrCreate(&user).
		Error

	return &user, err
}
