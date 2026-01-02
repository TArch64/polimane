package repository

import (
	"gorm.io/gorm"
)

func IncludeSoftDeleted(stmt *gorm.Statement) {
	stmt.Unscoped = true
}

var HardDelete = IncludeSoftDeleted
