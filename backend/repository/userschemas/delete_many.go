package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type DeleteManyOptions struct {
	UserID    model.ID
	SchemaIDs []model.ID
}

func (c *Client) DeleteMany(ctx context.Context, options *DeleteManyOptions) error {
	_, err := gorm.
		G[model.UserSchema](c.db).
		Where("user_id = ? AND schema_id IN (?)", options.UserID, options.SchemaIDs).
		Delete(ctx)

	return err
}
