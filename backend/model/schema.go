package model

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
)

const (
	SchemaPaletteSize = 9

	SchemaPatternSquare  = "square"
	SchemaPatternDiamond = "diamond"
)

type Schema struct {
	*Identifiable
	*Timestamps
	Name           string        `gorm:"not null;index;size:255" json:"name"`
	Palette        SchemaPalette `gorm:"not null;type:json" json:"palette,omitempty"`
	Content        SchemaContent `gorm:"not null;type:json" json:"content,omitempty"`
	ScreenshotedAt *time.Time    `json:"screenshotedAt"`
	Users          []User        `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade" json:"-"`
}

func SchemaScreenshotKey(id ID) string {
	return fmt.Sprintf("data/images/%s/schema.webp", id.String())
}

type SchemaPalette = datatypes.JSONSlice[string]
type SchemaContent = datatypes.JSONMap

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
