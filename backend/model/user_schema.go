package model

type UserSchema struct {
	*Timestamps
	UserID   ID          `gorm:"primaryKey" json:"userId"`
	SchemaID ID          `gorm:"primaryKey" json:"schemaId"`
	Access   AccessLevel `json:"access"`

	// Relations
	User   *User   `json:"-"`
	Schema *Schema `json:"-"`
}
