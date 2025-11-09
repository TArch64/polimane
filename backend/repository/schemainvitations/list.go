package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) List(ctx context.Context, filters ...repository.Filter) ([]*model.SchemaInvitation, error) {
	return gorm.
		G[*model.SchemaInvitation](c.db).
		Scopes(FilterAvailable).
		Scopes(filters...).
		Find(ctx)
}
