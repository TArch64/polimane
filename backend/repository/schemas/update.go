package schemas

import (
	"context"

	"polimane/backend/model"
)

type UpdateOptions struct {
	SchemaID model.ID
	Updates  *model.Schema
}

func (c *Client) Update(ctx context.Context, options *UpdateOptions) (err error) {
	return c.db.
		WithContext(ctx).
		Model(&model.Schema{
			Identifiable: &model.Identifiable{
				ID: options.SchemaID,
			},
		}).
		Updates(options.Updates).
		Error
}
