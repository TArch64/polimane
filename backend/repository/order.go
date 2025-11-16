package repository

import (
	"strings"

	"gorm.io/gorm"
)

func Order(orders ...string) Scope {
	return func(stmt *gorm.Statement) {
		stmt.Order(strings.Join(orders, ", "))
	}
}
