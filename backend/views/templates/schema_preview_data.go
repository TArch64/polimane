package templates

import "polimane/backend/model"

type SchemaPreviewData struct {
	Width  uint
	Height uint

	OffsetX int
	OffsetY int

	BeadSize    uint8
	ShapeCenter uint8
	ShapeRadius uint8

	Beads           model.SchemaBeads
	BackgroundColor string
}

func NewSchemaPreviewData(schema *model.Schema) *SchemaPreviewData {
	const beadSize = 12

	size := schema.Size.Data()
	beads := schema.Beads.Data()

	return &SchemaPreviewData{
		Width:  (uint(size.Left) + uint(size.Right)) * beadSize,
		Height: (uint(size.Top) + uint(size.Bottom)) * beadSize,

		OffsetX: int(size.Left) * beadSize,
		OffsetY: int(size.Top) * beadSize,

		BeadSize:    beadSize,
		ShapeCenter: beadSize / 2,
		ShapeRadius: (beadSize / 2) - 1,

		Beads:           beads,
		BackgroundColor: schema.BackgroundColor,
	}
}
