package helpers

import (
	"strconv"
	"strings"

	"polimane/backend/views/templates"
)

type SchemaBead struct {
	Color   string
	OffsetX int
	OffsetY int
}

func beadsGrid(data *templates.SchemaPreviewData) chan *SchemaBead {
	bufferSize := min(len(data.Beads)/20, 100)
	ch := make(chan *SchemaBead, bufferSize)

	initialOffsetX := -data.SizeLeft * int(data.BeadSize)
	initialOffsetY := -data.SizeTop * int(data.BeadSize)

	go func() {
		defer close(ch)

		for coord, color := range data.Beads {
			parts := strings.SplitN(coord, ":", 2)
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			ch <- &SchemaBead{
				OffsetX: initialOffsetX + (x * int(data.BeadSize)),
				OffsetY: initialOffsetY + (y * int(data.BeadSize)),
				Color:   color,
			}
		}
	}()

	return ch
}
