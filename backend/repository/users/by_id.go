package repositoryusers

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	"polimane/backend/services/db"
)

func ByID(ctx context.Context, id modelbase.ID) (*model.User, error) {
	var user model.User
	err := db.Instance.WithContext(ctx).Take(&user, id).Error
	return &user, err
}
