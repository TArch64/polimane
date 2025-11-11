package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) List(ctx context.Context, scopes ...repository.Scope) ([]*model.SchemaInvitation, error) {
	return gorm.
		G[*model.SchemaInvitation](c.db).
		Scopes(FilterAvailable).
		Scopes(scopes...).
		Find(ctx)
}
