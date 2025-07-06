package model

import (
	"polimane/backend/model/modelbase"
)

type UserSchema struct {
	*modelbase.Timestamps
	UserID   modelbase.ID `gorm:"primaryKey;autoIncrement:false" json:"userId"`
	SchemaID modelbase.ID `gorm:"primaryKey;autoIncrement:false" json:"schemaId"`
}
