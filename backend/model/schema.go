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
	Content        TSchemaContent `gorm:"not null;type:json" json:"content,omitempty"`
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

type SchemaPalette []string
type TSchemaPalette = datatypes.JSONType[SchemaPalette]

type TSchemaContent = datatypes.JSONType[*SchemaContent]

type SchemaContent struct {
	Size  *SchemaContentSize `json:"size"`
	Beads map[string]string
}

type SchemaContentSize struct {
	Left   uint8 `json:"left"`
	Top    uint8 `json:"top"`
	Right  uint8 `json:"right"`
	Bottom uint8 `json:"bottom"`
}
