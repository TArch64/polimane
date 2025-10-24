package model

import "time"

type SchemaInvitation struct {
	Email     string      `gorm:"primaryKey" json:"email"`
	SchemaID  ID          `gorm:"primaryKey" json:"schemaId"`
	Access    AccessLevel `json:"access"`
	ExpiresAt time.Time   `json:"expiresAt"`

	// Relations
	Schema *Schema `json:"-"`
}
