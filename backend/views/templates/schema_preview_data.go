package templates

import "polimane/backend/model"

type SchemaPreviewData struct {
	Width  uint
	Height uint

	BeadSize  uint8
	ShapeSize uint8

	SizeTop    int
	SizeRight  int
	SizeBottom int
	SizeLeft   int

	Beads           model.SchemaBeads
	BackgroundColor string
}

func NewSchemaPreviewData(schema *model.Schema) *SchemaPreviewData {
	const beadSize = 12

	size := schema.Size.Data()
	beads := schema.Beads.Data()

	return &SchemaPreviewData{
		Width:  (uint(size.Left) + uint(size.Right) + 2) * beadSize,
		Height: (uint(size.Top) + uint(size.Bottom) + 2) * beadSize,

		BeadSize:  beadSize,
		ShapeSize: beadSize - 2,

		SizeTop:    -int(size.Top),
		SizeRight:  int(size.Right),
		SizeBottom: int(size.Bottom),
		SizeLeft:   -int(size.Left),

		Beads:           beads,
		BackgroundColor: schema.BackgroundColor,
	}
}
