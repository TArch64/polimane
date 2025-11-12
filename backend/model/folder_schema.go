package model

type FolderSchema struct {
	*Timestamps
	FolderID ID `gorm:"primaryKey" json:"folderId"`
	SchemaID ID `gorm:"primaryKey" json:"schemaId"`

	// Relations
	Folder *Folder `json:"-"`
	Schema *Schema `json:"-"`
}
