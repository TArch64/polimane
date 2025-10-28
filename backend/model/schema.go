package model

import (
	"fmt"
	"strconv"
	"time"

	t "gorm.io/datatypes"
)

const (
	SchemaPaletteSize = 10
)

type Schema struct {
	*Identifiable
	*Timestamps
	Name            string                    `json:"name"`
	Palette         t.JSONType[SchemaPalette] `json:"palette,omitempty"`
	Size            t.JSONType[*SchemaSize]   `json:"size,omitempty"`
	Beads           t.JSONType[SchemaBeads]   `json:"beads,omitempty"`
	BackgroundColor string                    `gorm:"default:#f8f8f8" json:"backgroundColor"`
	ScreenshotedAt  *time.Time                `json:"screenshotedAt"`

	// Relations
	Users       []User             `gorm:"many2many:user_schemas" json:"-"`
	Invitations []SchemaInvitation `json:"-"`
}

type SchemaWithAccess struct {
	Schema
	Access AccessLevel `json:"access"`
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

type SchemaBeads map[string]*SchemaBead

type SchemaSize struct {
	Left   uint16 `validate:"required,gte=0,lte=65535" json:"left"`
	Top    uint16 `validate:"required,gte=0,lte=65535" json:"top"`
	Right  uint16 `validate:"required,gte=0,lte=65535" json:"right"`
	Bottom uint16 `validate:"required,gte=0,lte=65535" json:"bottom"`
}
