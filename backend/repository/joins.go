package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Join(expr clause.Expr) Scope {
	return func(stmt *gorm.Statement) {
		stmt.Scopes(func(db *gorm.DB) *gorm.DB {
			return db.Joins(expr.SQL, expr.Vars...)
		})
	}
}
