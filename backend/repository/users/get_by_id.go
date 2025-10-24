package users

import (
	"context"

	"polimane/backend/model"
)

func (c *Client) GetByID(ctx context.Context, id model.ID) (*model.User, error) {
	var user model.User
	err := c.db.WithContext(ctx).Take(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
