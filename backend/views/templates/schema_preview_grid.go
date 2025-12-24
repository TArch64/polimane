package templates

import (
	"math"
)

const (
	schemaPreviewPadding   = 40
	schemaPreviewMinWidth  = 1000
	schemaPreviewMinHeight = schemaPreviewMinWidth * 9 / 16
)

type SchemaBead struct {
	Color  string
	Circle *SchemaBeadCircle
	Bugle  *SchemaBeadBugle
}

func (s *SchemaBead) shift(x, y uint) {
	if s.Circle != nil {
		s.Circle.shift(x, y)
	} else if s.Bugle != nil {
		s.Bugle.shift(x, y)
	}
}

type SchemaBeadCircle struct {
	CenterX uint
	CenterY uint
}

func (s *SchemaBeadCircle) getShape(radius uint8) (x1, y1, x2, y2 uint) {
	x1 = s.CenterX - uint(radius)
	y1 = s.CenterY - uint(radius)
	x2 = s.CenterX + uint(radius)
	y2 = s.CenterY + uint(radius)
	return
}

func (s *SchemaBeadCircle) shift(x, y uint) {
	s.CenterX -= x
	s.CenterY -= y
}

type SchemaBeadBugle struct {
	X      uint
	Y      uint
	Width  uint
	Height uint
}

func (s *SchemaBeadBugle) getShape() (x1, y1, x2, y2 uint) {
	x1 = s.X
	y1 = s.Y
	x2 = x1 + s.Width
	y2 = y1 + s.Height
	return
}

func (s *SchemaBeadBugle) shift(x, y uint) {
	s.X -= x
	s.Y -= y
}

func (d *SchemaPreviewData) GenerateBeadsGrid() {
	for coord, bead := range d.Beads {
		if bead.Ref != nil {
			continue
		}

		x, y := coord.MustParseInt()

		item := SchemaBead{
			Color: bead.GetColor(),
		}

		var radialShiftX uint8
		if d.IsRadial && y%2 == 0 {
			radialShiftX = d.BeadSize / 2
		}

		if bead.Circle != nil {
			item.Circle = &SchemaBeadCircle{
				CenterX: uint(d.OffsetX+(x*int(d.BeadSize))) + uint(d.CircleCenter) + uint(radialShiftX),
				CenterY: uint(d.OffsetY+(y*int(d.BeadSize))) + uint(d.CircleCenter),
			}
			d.trySetShape(item.Circle.getShape(d.CircleCenter))
		} else if bead.Bugle != nil {
			spanX := x + int(bead.Bugle.Span.X)
			spanY := y + int(bead.Bugle.Span.Y)

			startCoordX := min(x, spanX)
			startCoordY := min(y, spanY)
			coordWidth := uint(math.Abs(float64(bead.Bugle.Span.X)) + 1)
			coordHeight := uint(math.Abs(float64(bead.Bugle.Span.Y)) + 1)

			item.Bugle = &SchemaBeadBugle{
				X:      uint(d.OffsetX+(startCoordX*int(d.BeadSize))) + uint(d.BuglePadding) + uint(radialShiftX),
				Y:      uint(d.OffsetY+(startCoordY*int(d.BeadSize))) + uint(d.BuglePadding),
				Width:  coordWidth*uint(d.BeadSize) - uint(d.BuglePadding)*2,
				Height: coordHeight*uint(d.BeadSize) - uint(d.BuglePadding)*2,
			}
			d.trySetShape(item.Bugle.getShape())
		} else {
			continue
		}

		d.BeadsGrid = append(d.BeadsGrid, &item)
	}

	originalWidth := d.Width - (d.MinX - schemaPreviewPadding*2)
	originalHeight := d.Height - (d.MinY - schemaPreviewPadding*2)

	d.Width = max(originalWidth, schemaPreviewMinWidth)
	d.Height = max(originalHeight, schemaPreviewMinHeight)

	paddingX := (d.Width - originalWidth) / 2
	paddingY := (d.Height - originalHeight) / 2

	for _, bead := range d.BeadsGrid {
		bead.shift(d.MinX-paddingX-schemaPreviewPadding, d.MinY-paddingY-schemaPreviewPadding)
	}
}

func (d *SchemaPreviewData) trySetShape(x, y, width, height uint) {
	if d.MinX == 0 || x < d.MinX {
		d.MinX = x
	}
	if d.MinY == 0 || y < d.MinY {
		d.MinY = y
	}
	if d.Width == 0 || width > d.Width {
		d.Width = width
	}
	if d.Height == 0 || height > d.Height {
		d.Height = height
	}
}
