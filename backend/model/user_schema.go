package model

type UserSchema struct {
	*Timestamps
	*SoftDeletable
	UserID   ID          `gorm:"primaryKey" json:"userId"`
	SchemaID ID          `gorm:"primaryKey" json:"schemaId"`
	FolderID *ID         `json:"folderId"`
	Access   AccessLevel `json:"access"`

	// Relations
	User   *User   `json:"-"`
	Schema *Schema `json:"-"`
	Folder *Folder `json:"-"`
}
