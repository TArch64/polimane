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

func (i *Impl) Update(ctx context.Context, options *UpdateOptions) (err error) {
	if options.User != nil {
		err = i.userSchemas.HasAccess(ctx, options.User.ID, options.SchemaID)
		if err != nil {
			return err
		}
	}

	return i.db.
		WithContext(ctx).
		Model(&model.Schema{
			Identifiable: &model.Identifiable{
				ID: options.SchemaID,
			},
		}).
		Updates(options.Updates).
		Error
}
