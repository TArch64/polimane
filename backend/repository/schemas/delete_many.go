package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type DeleteOptions struct {
	SchemaIDs []model.ID
}

func (c *Client) DeleteMany(ctx context.Context, options *DeleteOptions) error {
	_, err := gorm.
		G[model.Schema](c.db).
		Where("id IN (?)", options.SchemaIDs).
		Delete(ctx)

	return err
}
