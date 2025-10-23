package schemas

import (
	"context"

	"polimane/backend/model"
)

type ByIDOptions struct {
	User     *model.User
	SchemaID model.ID
	Select   []string
}

func (c *Client) GetOutByID(ctx context.Context, options *ByIDOptions, out interface{}) error {
	query := c.db.WithContext(ctx).
		Table("schemas")

	if options.User != nil {
		query = query.Scopes(IncludeUserSchemaScope(options.User.ID))
	}

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	return query.Take(out, options.SchemaID).Error
}

func (c *Client) GetByID(ctx context.Context, options *ByIDOptions) (*model.Schema, error) {
	var schema model.Schema
	err := c.GetOutByID(ctx, options, &schema)
	return &schema, err
}
