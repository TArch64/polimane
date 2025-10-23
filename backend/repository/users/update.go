package users

import (
	"context"

	"polimane/backend/model"
)

type UpdateOptions struct {
	UserID  model.ID
	Updates *model.User
}

func (i *Impl) Update(ctx context.Context, options *UpdateOptions) error {
	return i.db.
		WithContext(ctx).
		Model(&model.User{
			Identifiable: &model.Identifiable{
				ID: options.UserID,
			},
		}).
		Updates(options.Updates).
		Error
}
