package schemainvitations

import (
	"gorm.io/gorm"

	"polimane/backend/repository"
)

func FilterAvailable(db *gorm.Statement) {
	repository.AddWhere(db, gorm.Expr("expires_at > now()"))
}

func FilterExpired(db *gorm.Statement) {
	repository.AddWhere(db, gorm.Expr("expires_at <= now()"))
}
