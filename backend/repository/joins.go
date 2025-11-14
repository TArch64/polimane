package repository

import "gorm.io/gorm"

func AddJoin(db *gorm.Statement, query string, args ...interface{}) {
	db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	})
}
