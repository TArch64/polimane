package usercreate

import (
	"context"
	"errors"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryusers "polimane/backend/repository/users"
)

func (s *Service) GetOrCreate(ctx context.Context, workosUser *usermanagement.User) (*model.User, error) {
	user, err := s.users.Get(ctx,
		repositoryusers.WorkosIDEq(workosUser.ID),
		repository.IncludeSoftDeleted,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return s.Create(ctx, workosUser)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
