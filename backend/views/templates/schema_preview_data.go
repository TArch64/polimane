package templates

import "polimane/backend/model"

type SchemaPreviewData struct {
	MinX   uint
	MinY   uint
	Width  uint
	Height uint

	OffsetX int
	OffsetY int

	BeadSize          uint8
	CircleCenter      uint8
	CircleRadius      uint8
	BuglePadding      uint8
	BugleCornerRadius uint8

	Beads           model.SchemaBeads
	BackgroundColor string
	IsRadial        bool

	BeadsGrid []*SchemaBead
}

func NewSchemaPreviewData(schema *model.Schema) *SchemaPreviewData {
	const beadSize = 12
	const circleCenter = beadSize / 2

	size := schema.Size.Data()
	beads := schema.Beads.Data()

	data := &SchemaPreviewData{
		OffsetX: int(size.Left) * beadSize,
		OffsetY: int(size.Top) * beadSize,

		BeadSize:          beadSize,
		CircleCenter:      circleCenter,
		CircleRadius:      circleCenter - 1,
		BuglePadding:      beadSize / 4,
		BugleCornerRadius: beadSize / 6,

		Beads:           beads,
		BackgroundColor: schema.BackgroundColor,
		IsRadial:        schema.Layout == model.SchemaRadial,
	}

	data.GenerateBeadsGrid()
	return data
}
