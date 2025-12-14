package repository

import (
	"strings"

	"gorm.io/gorm"
)

func Group(columns ...string) Scope {
	return func(stmt *gorm.Statement) {
		stmt.Group(strings.Join(columns, ", "))
	}
}
