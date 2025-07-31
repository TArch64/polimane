package schemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

type UpdateOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID modelbase.ID
	Updates  *model.Schema
}

func (c *Impl) Update(options *UpdateOptions) (err error) {
	err = c.userSchemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return err
	}

	return c.db.
		WithContext(options.Ctx).
		Model(&model.Schema{
			Identifiable: &modelbase.Identifiable{
				ID: options.SchemaID,
			},
		}).
		Updates(options.Updates).
		Error
}
