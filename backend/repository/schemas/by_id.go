package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/db"
)

type ByIDOptions struct {
	Ctx      context.Context
	User     *model.User
	SchemaID modelbase.ID
	Select   []string
}

func ByID(options *ByIDOptions) (*model.Schema, error) {
	var err error

	err = repositoryuserschemas.HasAccess(options.Ctx, options.User.ID, options.SchemaID)
	if err != nil {
		return nil, err
	}

	var schema model.Schema
	query := db.Client().WithContext(options.Ctx)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	err = query.Take(&schema, options.SchemaID).Error
	return &schema, err
}
