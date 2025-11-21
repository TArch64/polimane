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

func IncludeSchemasScope(conditions ...clause.Expr) repository.Scope {
	return func(stmt *gorm.Statement) {
		expr := gorm.Expr("JOIN schemas ON schemas.id = user_schemas.schema_id")
		for _, condition := range conditions {
			expr.SQL += " AND " + condition.SQL
			expr.Vars = append(expr.Vars, condition.Vars...)
		}

		repository.AddJoin(stmt, expr)
	}
}

func FolderIDEq(id *model.ID) repository.Scope {
	return func(stmt *gorm.Statement) {
		var expr clause.Expr
		if id == nil {
			expr = gorm.Expr("folder_id IS NULL")
		} else {
			expr = gorm.Expr("folder_id = ?", *id)
		}
		repository.AddWhere(stmt, expr)
	}
}
