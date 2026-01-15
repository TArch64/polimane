package usercreate

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
)

func (s *Service) createUser(ctx context.Context, tx *gorm.DB, workosUser *usermanagement.User) (*model.User, error) {
	user := &model.User{
		WorkosID:  workosUser.ID,
		Email:     workosUser.Email,
		FirstName: workosUser.FirstName,
		LastName:  workosUser.LastName,
	}

	if err := s.users.InsertTx(ctx, tx, user); err != nil {
		return nil, err
	}

	return user, nil
}
