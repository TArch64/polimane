package repository

import (
	"gorm.io/gorm"
)

var (
	First = Limit(1)
)

func Limit(limit uint8) Scope {
	return func(stmt *gorm.Statement) {
		stmt.Limit(int(limit))
	}
}
