package users

import (
	"gorm.io/gorm"

	"polimane/backend/repository"
)

func WorkosIDEq(id string) repository.Scope {
	return func(stmt *gorm.Statement) {
		repository.AddWhere(stmt, gorm.Expr("workos_id = ?", id))
	}
}
