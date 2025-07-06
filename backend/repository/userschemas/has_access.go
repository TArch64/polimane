package repositoryuserschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	"polimane/backend/services/db"
)

func HasAccess(ctx context.Context, userID, schemaID modelbase.ID) error {
	var exists bool

	err := db.Client().
		WithContext(ctx).
		Model(&model.UserSchema{}).
		Select("1 AS exists").
		Where("user_id = ? AND schema_id = ?", userID, schemaID).
		Pluck("exists", &exists).
		Error

	if err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return nil
}
