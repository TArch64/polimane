package model

import (
	"fmt"
	"strconv"
	"time"

	t "gorm.io/datatypes"
)

type SchemaLayout string

const (
	SchemaPaletteSize = 10

	SchemaLinear SchemaLayout = "linear"
	SchemaRadial SchemaLayout = "radial"
)

type Schema struct {
	*Identifiable
	*Timestamps
	*SoftDeletable
	Name            string                    `json:"name"`
	Palette         t.JSONType[SchemaPalette] `json:"palette,omitempty"`
	Size            t.JSONType[*SchemaSize]   `json:"size,omitempty"`
	Beads           t.JSONType[SchemaBeads]   `json:"beads,omitempty"`
	BackgroundColor string                    `gorm:"default:#f8f8f8" json:"backgroundColor"`
	Layout          SchemaLayout              `json:"layout"`
	ScreenshotedAt  *time.Time                `json:"screenshotedAt"`

	// Relations
	Users       []User             `gorm:"many2many:user_schemas" json:"-"`
	Folders     []Folder           `gorm:"many2many:folder_schemas" json:"-"`
	Invitations []SchemaInvitation `json:"-"`
}

func SchemaScreenshotPath(id ID, timestamp *time.Time) *string {
	if timestamp == nil {
		return nil
	}
	path := schemaScreenshotPath(id)
	path += "?v="
	path += strconv.FormatInt(timestamp.Unix(), 10)
	return &path
}

func schemaScreenshotPath(id ID) string {
	return fmt.Sprintf("images/%s/schema.svg", id.String())
}

func SchemaScreenshotKey(id ID) string {
	return "data/" + schemaScreenshotPath(id)
}

type SchemaPalette []string

type SchemaBeads map[SchemaCoord]*SchemaBead

type SchemaSize struct {
	Left   uint16 `validate:"required,gte=0,lte=65535" json:"left"`
	Top    uint16 `validate:"required,gte=0,lte=65535" json:"top"`
	Right  uint16 `validate:"required,gte=0,lte=65535" json:"right"`
	Bottom uint16 `validate:"required,gte=0,lte=65535" json:"bottom"`
}
