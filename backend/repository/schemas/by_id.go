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

func (i *Client) OutByID(ctx context.Context, options *ByIDOptions, out interface{}) error {
	query := i.db.WithContext(ctx).
		Table("schemas")

	if options.User != nil {
		query = query.Scopes(UserSchemaScope(options.User.ID))
	}

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	return query.Take(out, options.SchemaID).Error
}

func (i *Client) ByID(ctx context.Context, options *ByIDOptions) (*model.Schema, error) {
	var schema model.Schema
	err := i.OutByID(ctx, options, &schema)
	return &schema, err
}
