package repository

import "gorm.io/gorm"

func AddJoin(stmt *gorm.Statement, query string, args ...interface{}) {
	stmt.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	})
}
