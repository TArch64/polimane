package model

import (
	t "gorm.io/datatypes"
)

type UserSchema struct {
	*Timestamps
	*SoftDeletable
	UserID   ID                          `gorm:"primaryKey" json:"userId"`
	SchemaID ID                          `gorm:"primaryKey" json:"schemaId"`
	FolderID *ID                         `json:"folderId"`
	Access   AccessLevel                 `json:"access"`
	Counters t.JSONType[*SchemaCounters] `json:"counters" gorm:"default:'{}'::json"`

	// Relations
	User   *User   `json:"-"`
	Schema *Schema `json:"-"`
	Folder *Folder `json:"-"`
}

type SchemaCounters struct {
	SchemaBeads  uint16 `json:"schemaBeads"`
	SharedAccess uint8  `json:"sharedAccess"`
}
