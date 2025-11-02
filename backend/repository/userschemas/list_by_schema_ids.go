package userschemas

import (
	"context"
	"strings"

	"polimane/backend/model"
)

type ListBySchemaIDsOptions struct {
	SchemaIDs []model.ID
	Scopes    []model.Scope
	Select    []string
	Order     []string
}

func (c *Client) ListBySchemaIDsOut(ctx context.Context, options *ListBySchemaIDsOptions, out interface{}) error {
	query := c.db.
		WithContext(ctx).
		Table("user_schemas").
		Where("user_schemas.schema_id IN (?)", options.SchemaIDs)

	if len(options.Scopes) > 0 {
		query.Scopes(options.Scopes...)
	}

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	if len(options.Order) > 0 {
		query = query.Order(strings.Join(options.Order, ", "))
	}

	return query.Limit(100).Find(out).Error
}
