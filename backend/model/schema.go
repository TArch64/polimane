package model

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/datatypes"
)

const (
	SchemaPaletteSize = 9
)

type Schema struct {
	*Identifiable
	*Timestamps
	Name           string         `gorm:"not null;index;size:255" json:"name"`
	Palette        TSchemaPalette `gorm:"not null;type:json" json:"palette,omitempty"`
	Size           TSchemaSize    `gorm:"not null;type:json" json:"size,omitempty"`
	Beads          TSchemaBeads   `gorm:"not null;type:json" json:"beads,omitempty"`
	ScreenshotedAt *time.Time     `json:"screenshotedAt"`
	Users          []User         `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade" json:"-"`
}

func (s *Schema) ScreenshotPath() *string {
	if s.ScreenshotedAt == nil {
		return nil
	}
	path := SchemaScreenshotKey(s.ID)
	path += "?v=" + strconv.FormatInt(s.ScreenshotedAt.Unix(), 10)
	return &path
}

func SchemaScreenshotKey(id ID) string {
	return fmt.Sprintf("data/images/%s/schema.webp", id.String())
}

type TSchemaPalette = datatypes.JSONType[SchemaPalette]
type TSchemaSize = datatypes.JSONType[*SchemaSize]
type TSchemaBeads = datatypes.JSONType[SchemaBeads]

type SchemaPalette []string

type SchemaBeads map[string]string

type SchemaSize struct {
	Left   uint8 `validate:"required,gte=0,lte=255" json:"left"`
	Top    uint8 `validate:"required,gte=0,lte=255" json:"top"`
	Right  uint8 `validate:"required,gte=0,lte=255" json:"right"`
	Bottom uint8 `validate:"required,gte=0,lte=255" json:"bottom"`
}
