package model

import (
	"polimane/backend/model/modelbase"
)

type User struct {
	*modelbase.Identifiable
	WorkosID string    `gorm:"not null;uniqueIndex;size:32" json:"-"`
	Schemas  []*Schema `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade;" json:"-"`
}
