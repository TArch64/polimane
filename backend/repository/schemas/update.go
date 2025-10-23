package schemas

import (
	"context"

	"polimane/backend/model"
)

type UpdateOptions struct {
	User     *model.User
	SchemaID model.ID
	Updates  *model.Schema
}

func (c *Client) Update(ctx context.Context, options *UpdateOptions) (err error) {
	if options.User != nil {
		err = c.userSchemas.HasAccess(ctx, options.User.ID, options.SchemaID, model.AccessWrite)
		if err != nil {
			return err
		}
	}

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
