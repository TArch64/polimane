package schemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUserSchemaScope(userID model.ID) repository.Scope {
	return func(db *gorm.Statement) {
		repository.AddJoin(db, "JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
	}
}

func FilterByFolder(id *model.ID) repository.Scope {
	if id == nil {
		return func(db *gorm.Statement) {
			repository.AddJoin(db, "LEFT JOIN folder_schemas ON schemas.id = folder_schemas.schema_id")
			repository.AddWhere(db, gorm.Expr("folder_schemas.schema_id IS NULL"))
		}
	} else {
		return func(db *gorm.Statement) {}
	}
}
