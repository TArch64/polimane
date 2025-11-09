package users

import (
	"gorm.io/gorm"

	"polimane/backend/repository"
)

func WorkosIDEq(id string) repository.Scope {
	return func(db *gorm.Statement) {
		repository.AddWhere(db, gorm.Expr("workos_id = ?", id))
	}
}
