package repository

import (
	"gorm.io/gorm"
)

func Join(query string, args ...interface{}) Scope {
	return func(stmt *gorm.Statement) {
		stmt.Scopes(func(db *gorm.DB) *gorm.DB {
			return db.Joins(query, args...)
		})
	}
}

func Preload(query string, args ...interface{}) Scope {
	return func(stmt *gorm.Statement) {
		stmt.Preload(query, args...)
	}
}
