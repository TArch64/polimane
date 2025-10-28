package schemas

import (
	"context"

	"polimane/backend/model"
)

type CountByUserOptions struct {
	User *model.User
}

func (c *Client) CountByUserOut(ctx context.Context, options *CountByUserOptions, out *int64) error {
	return c.db.
		WithContext(ctx).
		Table("schemas").
		Scopes(IncludeUserSchemaScope(options.User.ID)).
		Count(out).
		Error
}
