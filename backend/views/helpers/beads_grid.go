package helpers

import (
	"strconv"
	"strings"

	"polimane/backend/views/templates"
)

type SchemaBead struct {
	Color   string
	CenterX int
	CenterY int
}

func beadsGrid(data *templates.SchemaPreviewData) chan *SchemaBead {
	bufferSize := min(len(data.Beads)/20, 100)
	ch := make(chan *SchemaBead, bufferSize)

	go func() {
		defer close(ch)

		for coord, bead := range data.Beads {
			parts := strings.SplitN(coord, ":", 2)
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			ch <- &SchemaBead{
				CenterX: data.OffsetX + (x * int(data.BeadSize)) + int(data.ShapeCenter),
				CenterY: data.OffsetY + (y * int(data.BeadSize)) + int(data.ShapeCenter),
				Color:   bead.Circle.Color,
			}
		}
	}()

	return ch
}
