package schemas

import (
	"context"

	"polimane/backend/model"
)

type UpdateOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID model.ID
	Updates  *model.Schema
}

func (c *Client) Update(options *UpdateOptions) (err error) {
	err = c.userSchemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return err
	}

	return c.db.
		WithContext(options.Ctx).
		Model(&model.Schema{
			Identifiable: &model.Identifiable{
				ID: options.SchemaID,
			},
		}).
		Updates(options.Updates).
		Error
}
