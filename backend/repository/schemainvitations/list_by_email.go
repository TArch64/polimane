package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) ListByEmail(ctx context.Context, email string) ([]*model.SchemaInvitation, error) {
	return gorm.
		G[*model.SchemaInvitation](c.db).
		Where("email = ? and expires_at > now()", email).
		Find(ctx)
}
