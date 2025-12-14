package model

import (
	"gorm.io/gorm"
)

type SoftDeletable struct {
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
