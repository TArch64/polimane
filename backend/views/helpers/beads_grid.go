package helpers

import (
	"math"
	"strconv"
	"strings"

	"polimane/backend/views/templates"
)

type SchemaBead struct {
	Color  string
	Circle *SchemaBeadCircle
	Bugle  *SchemaBeadBugle
}

type SchemaBeadCircle struct {
	CenterX int
	CenterY int
}

type SchemaBeadBugle struct {
	X      int
	Y      int
	Width  int
	Height int
}

func beadsGrid(data *templates.SchemaPreviewData) chan *SchemaBead {
	bufferSize := min(len(data.Beads)/20, 100)
	ch := make(chan *SchemaBead, bufferSize)

	go func() {
		defer close(ch)

		for coord, bead := range data.Beads {
			if bead.Ref != nil {
				continue
			}

			parts := strings.SplitN(coord, ":", 2)
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			item := SchemaBead{
				Color: bead.GetColor(),
			}

			if bead.Circle != nil {
				item.Circle = &SchemaBeadCircle{
					CenterX: data.OffsetX + (x * int(data.BeadSize)) + int(data.CircleCenter),
					CenterY: data.OffsetY + (y * int(data.BeadSize)) + int(data.CircleCenter),
				}
			} else if bead.Bugle != nil {
				spanX := x + int(bead.Bugle.Span.X)
				spanY := y + int(bead.Bugle.Span.Y)

				startCoordX := min(x, spanX)
				startCoordY := min(y, spanY)
				coordWidth := int(math.Abs(float64(bead.Bugle.Span.X)) + 1)
				coordHeight := int(math.Abs(float64(bead.Bugle.Span.Y)) + 1)

				item.Bugle = &SchemaBeadBugle{
					X:      data.OffsetX + (startCoordX * int(data.BeadSize)) + int(data.BuglePadding),
					Y:      data.OffsetY + (startCoordY * int(data.BeadSize)) + int(data.BuglePadding),
					Width:  coordWidth*int(data.BeadSize) - int(data.BuglePadding)*2,
					Height: coordHeight*int(data.BeadSize) - int(data.BuglePadding)*2,
				}
			} else {
				continue
			}

			ch <- &item
		}
	}()

	return ch
}
