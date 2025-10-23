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

func (i *Impl) ByID(ctx context.Context, options *ByIDOptions) (*model.Schema, error) {
	var err error

	if options.User != nil {
		err = i.userSchemas.HasAccess(ctx, options.User.ID, options.SchemaID)
		if err != nil {
			return nil, err
		}
	}

	query := i.db.WithContext(ctx)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	var schema model.Schema
	err = query.Take(&schema, options.SchemaID).Error
	return &schema, err
}
