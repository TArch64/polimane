package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddJoin(stmt *gorm.Statement, expr clause.Expr) {
	stmt.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Joins(expr.SQL, expr.Vars...)
	})
}
