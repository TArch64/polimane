package model

type UserSchema struct {
	*Timestamps
	UserID   ID `gorm:"type:uuid;primaryKey" json:"userId"`
	SchemaID ID `gorm:"type:uuid;primaryKey" json:"schemaId"`
}
