package userschemas

import (
	"context"

	"polimane/backend/model"
)

type ListBySchemaIDOptions struct {
	SchemaID model.ID
	Select   []string
}

func (c *Client) ListBySchemaIDOut(ctx context.Context, options *ListBySchemaIDOptions, out interface{}) error {
	query := c.db.
		WithContext(ctx).
		Table("user_schemas").
		Joins("JOIN users ON users.id = user_schemas.user_id AND user_schemas.schema_id = ?", options.SchemaID)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	return query.
		Limit(100).
		Order("user_schemas.created_at ASC").
		Find(out).
		Error
}
