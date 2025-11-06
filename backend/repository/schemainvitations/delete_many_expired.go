package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) DeleteManyExpired(ctx context.Context) error {
	_, err := gorm.
		G[model.SchemaInvitation](c.db).
		Where("expires_at < now()").
		Delete(ctx)

	return err
}
