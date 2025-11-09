package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) ListSchemasAccessOut(ctx context.Context, schemaIDs, out interface{}) error {
	return gorm.
		G[*model.SchemaInvitation](c.db).
		Select("email, MIN(access) AS access, MIN(access) != MAX(access) as is_uneven_access").
		Where("schema_id IN (?)", schemaIDs).
		Group("email").
		Order("email").
		Scan(ctx, out)
}
