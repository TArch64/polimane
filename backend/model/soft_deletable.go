package model

import (
	"time"

	"gorm.io/gorm"
)

type SoftDeletable struct {
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func SoftDeletedNow() *SoftDeletable {
	return &SoftDeletable{
		DeletedAt: gorm.DeletedAt{
			Valid: true,
			Time:  time.Now(),
		},
	}
}
