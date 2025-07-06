package repositoryusers

import (
	"context"

	"polimane/backend/model"
	"polimane/backend/services/db"
)

func ByName(ctx context.Context, username string) (*model.User, error) {
	user := model.User{Name: username}
	err := db.Client().WithContext(ctx).Where(&user).Take(&user).Error
	return &user, err
}
