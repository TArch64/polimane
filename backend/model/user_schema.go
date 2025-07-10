package model

import (
	"polimane/backend/model/modelbase"
)

type UserSchema struct {
	*modelbase.Timestamps
	UserID   modelbase.ID `gorm:"type:uuid;primaryKey" json:"userId"`
	SchemaID modelbase.ID `gorm:"type:uuid;primaryKey" json:"schemaId"`
}
