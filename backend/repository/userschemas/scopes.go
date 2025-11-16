package userschemas

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUsersLegacyScope(db *gorm.DB) *gorm.DB {
	return db.Joins("JOIN users ON users.id = user_schemas.user_id")
}

func IncludeSchemasScope(stmt *gorm.Statement) {
	expr := gorm.Expr("JOIN schemas ON schemas.id = user_schemas.schema_id")
	repository.AddJoin(stmt, expr)
}

func FolderIDEq(id *model.ID) repository.Scope {
	var expr clause.Expr
	if id == nil {
		expr = gorm.Expr("folder_id IS NULL")
	} else {
		expr = gorm.Expr("folder_id = ?", *id)
	}

	return func(stmt *gorm.Statement) {
		repository.AddWhere(stmt, expr)
	}
}
