package schemas

import (
	"context"

	"polimane/backend/model"
)

type CountByUserOptions struct {
	User *model.User
}

func (c *Client) CountByUser(ctx context.Context, options *CountByUserOptions) (int64, error) {
	var count int64
	err := c.db.
		WithContext(ctx).
		Table("schemas").
		Scopes(IncludeUserSchemaScope(options.User.ID)).
		Count(&count).
		Error

	return count, err
}
