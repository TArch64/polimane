package schemas

import (
	"context"

	"polimane/backend/model"
)

type ByUserOptions struct {
	User   *model.User
	Select []string
}

func (i *Client) OutByUser(ctx context.Context, options *ByUserOptions, out interface{}) error {
	query := i.db.
		WithContext(ctx).
		Table("schemas").
		Scopes(UserSchemaScope(options.User.ID))

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	return query.
		Limit(100).
		Order("schemas.created_at DESC").
		Find(out).
		Error
}

func (i *Client) ByUser(ctx context.Context, options *ByUserOptions) ([]*model.Schema, error) {
	var schemas []*model.Schema
	err := i.OutByUser(ctx, options, &schemas)
	return schemas, err
}
