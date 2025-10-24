package schemainvitations

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type ListBySchemaIDOptions struct {
	SchemaID model.ID
	Select   []string
}

func (c *Client) ListBySchemaIDOut(ctx context.Context, options *ListBySchemaIDOptions, out interface{}) error {
	query := gorm.
		G[*model.SchemaInvitation](c.db).
		Where("schema_id = ?", options.SchemaID)

	if len(options.Select) > 0 {
		query = query.Select(strings.Join(options.Select, ", "))
	}

	return query.Scan(ctx, out)
}
