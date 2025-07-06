package repositoryschemas

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/services/db"
)

type ByUserOptions struct {
	Ctx        context.Context
	User       *model.User
	Attributes []string
}

func ByUser(options *ByUserOptions) ([]*model.Schema, error) {
	var schemas []*model.Schema

	err := db.Client().
		Joins("JOIN user_schemas ON user_schemas.user_id = ?", options.User.ID).
		Find(&schemas).
		Error

	return schemas, err
}
