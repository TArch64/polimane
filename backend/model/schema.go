package model

import (
	"fmt"
	"strconv"
	"time"

	t "gorm.io/datatypes"
)

const (
	SchemaPaletteSize = 9
)

type Schema struct {
	*Identifiable
	*Timestamps
	Name            string                    `gorm:"not null;index;size:255" json:"name"`
	Palette         t.JSONType[SchemaPalette] `gorm:"not null;type:json" json:"palette,omitempty"`
	Size            t.JSONType[*SchemaSize]   `gorm:"not null;type:json" json:"size,omitempty"`
	Beads           t.JSONType[SchemaBeads]   `gorm:"not null;type:json" json:"beads,omitempty"`
	BackgroundColor string                    `gorm:"not null;size:30;default:#f8f8f8" json:"backgroundColor"`
	ScreenshotedAt  *time.Time                `json:"screenshotedAt"`
	Users           []User                    `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade" json:"-"`
}

func (s *Schema) ScreenshotPath() *string {
	if s.ScreenshotedAt == nil {
		return nil
	}
	path := schemaScreenshotPath(s.ID)
	path += "?v="
	path += strconv.FormatInt(s.ScreenshotedAt.Unix(), 10)
	return &path
}

func (s *Schema) ScreenshotKey() string {
	return SchemaScreenshotKey(s.ID)
}

func schemaScreenshotPath(id ID) string {
	return fmt.Sprintf("images/%s/schema.svg", id.String())
}

func SchemaScreenshotKey(id ID) string {
	return "data/" + schemaScreenshotPath(id)
}

type SchemaPalette []string

type SchemaBeads map[string]string

type SchemaSize struct {
	Left   uint8 `validate:"required,gte=0,lte=255" json:"left"`
	Top    uint8 `validate:"required,gte=0,lte=255" json:"top"`
	Right  uint8 `validate:"required,gte=0,lte=255" json:"right"`
	Bottom uint8 `validate:"required,gte=0,lte=255" json:"bottom"`
}
