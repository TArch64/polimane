package schemainvitations

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type ListBySchemaIDsOptions struct {
	SchemaIDs []model.ID
	Select    []string
	Order     []string
}

func (c *Client) ListBySchemaIDsOut(ctx context.Context, options *ListBySchemaIDsOptions, out interface{}) error {
	query := gorm.
		G[*model.SchemaInvitation](c.db).
		Where("schema_id IN (?)", options.SchemaIDs)

	if len(options.Select) > 0 {
		query = query.Select(strings.Join(options.Select, ", "))
	}

	if len(options.Order) > 0 {
		query = query.Order(strings.Join(options.Order, ", "))
	}

	return query.Scan(ctx, out)
}
