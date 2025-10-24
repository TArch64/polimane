package schemas

import (
	"context"

	"polimane/backend/model"
)

type ByUserOptions struct {
	User   *model.User
	Select []string
}

func (c *Client) ListByUserOut(ctx context.Context, options *ByUserOptions, out interface{}) error {
	query := c.db.
		WithContext(ctx).
		Table("schemas").
		Scopes(IncludeUserSchemaScope(options.User.ID))

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	return query.
		Limit(100).
		Order("schemas.created_at DESC").
		Find(out).
		Error
}

func (c *Client) ListByUser(ctx context.Context, options *ByUserOptions) ([]*model.Schema, error) {
	var schemas []*model.Schema
	err := c.ListByUserOut(ctx, options, &schemas)
	return schemas, err
}
