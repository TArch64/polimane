package users

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

func (c *Client) ByID(ctx context.Context, id modelbase.ID) (*model.User, error) {
	var user model.User
	err := c.db.WithContext(ctx).Take(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
