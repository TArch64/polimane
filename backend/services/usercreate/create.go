package usercreate

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/model"
)

func (s *Service) Create(ctx context.Context, workosUser *usermanagement.User) (user *model.User, err error) {
	err = s.users.DB.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			if user, err = s.createUser(ctx, tx, workosUser); err != nil {
				return err
			}

			if err = s.createSubscription(ctx, tx, user); err != nil {
				return err
			}

			return s.acceptInvitations(ctx, tx, user)
		})

	return
}
