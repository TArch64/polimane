package schemas

import (
	"context"

	"polimane/backend/model"
)

type ByUserOptions struct {
	User   *model.User
	Select []string
}

func (i *Impl) ByUser(ctx context.Context, options *ByUserOptions) ([]*model.Schema, error) {
	query := i.db.
		WithContext(ctx).
		Scopes(UserSchemaScope(options.User.ID))

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	var schemas []*model.Schema

	err := query.
		Limit(100).
		Order("schemas.created_at DESC").
		Find(&schemas).
		Error

	return schemas, err
}
