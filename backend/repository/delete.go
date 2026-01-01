package repository

import (
	"gorm.io/gorm"
)

func HardDelete(stmt *gorm.Statement) {
	stmt.Unscoped = true
}

func IncludeSoftDeleted(stmt *gorm.Statement) {
	stmt.Unscoped = true
}
