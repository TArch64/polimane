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

func (i *Impl) Update(options *UpdateOptions) (err error) {
	err = i.userSchemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return err
	}

	return i.db.
		WithContext(options.Ctx).
		Model(&model.Schema{
			Identifiable: &model.Identifiable{
				ID: options.SchemaID,
			},
		}).
		Updates(options.Updates).
		Error
}
