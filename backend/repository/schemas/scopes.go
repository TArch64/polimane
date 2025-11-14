package schemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUserSchemaLegacyScope(userID model.ID) model.LegacyScope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
	}
}

func IncludeUserSchemaScope(userID model.ID) repository.Scope {
	return func(db *gorm.Statement) {
		repository.AddJoin(db, "JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
	}
}
