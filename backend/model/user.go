package model

import (
	"polimane/backend/model/modelbase"
)

type User struct {
	*modelbase.Identifiable
	*modelbase.Timestamps
	Name         string    `gorm:"not null;uniqueIndex;size:255" json:"name"`
	PasswordHash string    `gorm:"not null;type:text" json:"-"`
	Schemas      []*Schema `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade;" json:"-"`
}
