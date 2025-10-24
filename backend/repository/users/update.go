package users

import (
	"context"

	"polimane/backend/model"
)

type UpdateOptions struct {
	UserID  model.ID
	Updates *model.User
}

func (c *Client) Update(ctx context.Context, options *UpdateOptions) error {
	return c.db.
		WithContext(ctx).
		Model(&model.User{
			Identifiable: &model.Identifiable{
				ID: options.UserID,
			},
		}).
		Updates(options.Updates).
		Error
}
