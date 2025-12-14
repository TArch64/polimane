package schemainvitations

import (
	"polimane/backend/repository"
)

var (
	FilterAvailable = repository.Where("expires_at > now()")
	FilterExpired   = repository.Where("expires_at <= now()")
)
