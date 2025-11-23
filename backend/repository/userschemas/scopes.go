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

var (
	FilterWithoutFolder = repository.Where(gorm.Expr("folder_id IS NULL"))
)

func FolderIDEq(id *model.ID) repository.Scope {
	if id == nil {
		return FilterWithoutFolder
	} else {
		return repository.Where(gorm.Expr("folder_id = ?", *id))
	}
}
