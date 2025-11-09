package schemas

import (
	"context"

	"polimane/backend/model"
)

type ListByUserOptions struct {
	User       *model.User
	Pagination *model.Pagination
	Select     []string
}

func (c *Client) ListByUserOut(ctx context.Context, options *ListByUserOptions, out interface{}) error {
	query := c.db.
		WithContext(ctx).
		Table("schemas").
		Scopes(
			IncludeUserSchemaScope(options.User.ID),
			model.PaginationScope(options.Pagination),
		)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	return query.
		Limit(model.DefaultBatch).
		Order("schemas.created_at DESC").
		Find(out).
		Error
}

func (c *Client) ListByUser(ctx context.Context, options *ListByUserOptions) ([]*model.Schema, error) {
	var schemas []*model.Schema
	err := c.ListByUserOut(ctx, options, &schemas)
	return schemas, err
}
