package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) ListSchemasAccessOut(ctx context.Context, schemaIDs []model.ID, out interface{}) error {
	return gorm.
		G[*model.SchemaInvitation](c.db).
		Select("email, MIN(access) AS access, MIN(access) != MAX(access) as is_uneven_access").
		Scopes(repository.InSchemaIDs(schemaIDs)).
		Group("email").
		Order("email").
		Scan(ctx, out)
}
