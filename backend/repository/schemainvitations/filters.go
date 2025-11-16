package schemainvitations

import (
	"gorm.io/gorm"

	"polimane/backend/repository"
)

func FilterAvailable(stmt *gorm.Statement) {
	repository.AddWhere(stmt, gorm.Expr("expires_at > now()"))
}

func FilterExpired(stmt *gorm.Statement) {
	repository.AddWhere(stmt, gorm.Expr("expires_at <= now()"))
}
