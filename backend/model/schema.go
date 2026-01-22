package model

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/datatypes"
)

type SchemaLayout string

const (
	SchemaLinear SchemaLayout = "linear"
	SchemaRadial SchemaLayout = "radial"
)

var (
	defaultSchemaWidth  uint16 = 50
	defaultSchemaHeight uint16 = 15

	DefaultSchemaSize = SchemaSize{
		Left:   defaultSchemaWidth,
		Right:  defaultSchemaWidth - 1,
		Top:    defaultSchemaHeight,
		Bottom: defaultSchemaHeight - 1,
	}

	DefaultPalette = SchemaPalette{
		"#333333",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
)

type Schema struct {
	*Identifiable
	*Timestamps
	*SoftDeletable
	Name            string            `json:"name"`
	Palette         SchemaPaletteJSON `json:"palette,omitempty"`
	Size            SchemaSizeJSON    `json:"size,omitempty"`
	Beads           SchemaBeadsJSON   `json:"beads,omitempty"`
	BackgroundColor string            `gorm:"default:#f8f8f8" json:"backgroundColor"`
	Layout          SchemaLayout      `json:"layout"`
	ScreenshotedAt  *time.Time        `json:"screenshotedAt"`
	DeletedBy       *ID               `json:"-"`

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
type SchemaPaletteJSON = datatypes.JSONType[SchemaPalette]

type SchemaSize struct {
	Left   uint16 `validate:"required,gte=0,lte=65535" json:"left"`
	Top    uint16 `validate:"required,gte=0,lte=65535" json:"top"`
	Right  uint16 `validate:"required,gte=0,lte=65535" json:"right"`
	Bottom uint16 `validate:"required,gte=0,lte=65535" json:"bottom"`
}

type SchemaSizeJSON = datatypes.JSONType[*SchemaSize]

type SchemaBeads map[SchemaCoord]*SchemaBead

func (s SchemaBeads) CountVisible() uint16 {
	beadsCount := uint16(len(s))
	for _, bead := range s {
		if bead.Ref != nil {
			beadsCount--
		}
	}
	return beadsCount
}

type SchemaBeadsJSON = datatypes.JSONType[SchemaBeads]
