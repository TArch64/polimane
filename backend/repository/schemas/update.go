package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/db"
)

type UpdateOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID modelbase.ID
	Updates  *model.Schema
}

func Update(options *UpdateOptions) (err error) {
	err = repositoryuserschemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return err
	}

	return db.Instance.
		WithContext(options.Ctx).
		Model(&model.Schema{
			Identifiable: &modelbase.Identifiable{
				ID: options.SchemaID,
			},
		}).
		Updates(options.Updates).
		Error
}
