package schemainvitations

import (
	"gorm.io/gorm"

	"polimane/backend/repository"
)

var (
	FilterAvailable = repository.Where(gorm.Expr("expires_at > now()"))
	FilterExpired   = repository.Where(gorm.Expr("expires_at <= now()"))
)
