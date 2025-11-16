package schemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUserSchemaScope(userID model.ID) repository.Scope {
	return func(stmt *gorm.Statement) {
		repository.AddJoin(stmt, "JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
	}
}

func FilterByFolder(userID model.ID, folderID *model.ID) repository.Scope {
	if folderID == nil {
		return func(stmt *gorm.Statement) {
			subquery := gorm.
				G[model.FolderSchema](stmt.DB).
				Select("1").
				Scopes(func(stmt *gorm.Statement) {
					repository.AddJoin(stmt, "JOIN folders ON folder_schemas.folder_id = folders.id AND folders.user_id = ?", userID)
				}).
				Where("folder_schemas.schema_id = schemas.id")

			repository.AddWhere(stmt, gorm.Expr("NOT EXISTS (?)", subquery))
		}
	} else {
		return func(stmt *gorm.Statement) {}
	}
}
