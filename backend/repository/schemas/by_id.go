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
	query := i.db.WithContext(ctx)

	if options.User != nil {
		query = query.Scopes(UserSchemaScope(options.User.ID))
	}

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	var schema model.Schema
	err = query.Take(&schema, options.SchemaID).Error
	return &schema, err
}
