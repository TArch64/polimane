package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/services/db"
)

type ByUserOptions struct {
	Ctx    context.Context
	User   *model.User
	Select []string
}

func ByUser(options *ByUserOptions) ([]*model.Schema, error) {
	var schemas []*model.Schema

	query := db.Client().
		WithContext(options.Ctx).
		Joins("JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", options.User.ID)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	err := query.Find(&schemas).Error
	return schemas, err
}
