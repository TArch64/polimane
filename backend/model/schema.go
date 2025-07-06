package model

import (
	. "gorm.io/datatypes"

	"polimane/backend/model/modelbase"
)

const (
	SchemaPaletteSize = 9

	SchemaPatternSquare  = "square"
	SchemaPatternDiamond = "diamond"
)

type Schema struct {
	*modelbase.Identifiable
	*modelbase.Timestamps
	Name    string        `gorm:"not null;index;size:255" json:"name"`
	Palette SchemaPalette `gorm:"not null;type:json" json:"palette,omitempty"`
	Content SchemaContent `gorm:"not null;type:json" json:"content,omitempty"`
	Users   []User        `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade" json:"-"`
}

type SchemaPalette = JSONSlice[string]
type SchemaContent = JSONSlice[*SchemaPattern]

type SchemaPattern struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Type    string       `json:"type"`
	Content []*SchemaRow `json:"content"`
}

type SchemaRow struct {
	ID      string       `json:"id"`
	Content []SchemaBead `json:"content"`
}

type SchemaBead struct {
	ID    string `json:"id"`
	Color string `json:"color"`
}
