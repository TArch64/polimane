package schemas

import (
	"context"

	"polimane/backend/model"
)

type ByUserOptions struct {
	Ctx    context.Context
	User   *model.User
	Select []string
}

func (i *Impl) ByUser(options *ByUserOptions) ([]*model.Schema, error) {
	query := i.db.
		WithContext(options.Ctx).
		Joins("JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", options.User.ID)

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	var schemas []*model.Schema
	err := query.Find(&schemas).Error
	return schemas, err
}
