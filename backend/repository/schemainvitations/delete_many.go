package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type DeleteManyOptions struct {
	Email     string
	SchemaIDs []model.ID
}

func (c *Client) DeleteMany(ctx context.Context, options *DeleteManyOptions) error {
	_, err := gorm.
		G[model.SchemaInvitation](c.db).
		Where("email = ? AND schema_id IN (?)", options.Email, options.SchemaIDs).
		Delete(ctx)

	return err
}
