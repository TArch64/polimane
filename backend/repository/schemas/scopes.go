package schemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
)

func UserSchemaScope(userID model.ID) model.Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
	}
}
