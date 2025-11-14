package repository

import (
	"strings"

	"gorm.io/gorm"
)

func Order(orders ...string) Scope {
	return func(db *gorm.Statement) {
		db.Order(strings.Join(orders, ", "))
	}
}
