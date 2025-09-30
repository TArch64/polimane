package helpers

import (
	"html/template"
	"strconv"
	"strings"

	"polimane/backend/views/templates"
)

type SchemaBead struct {
	Color     string
	OffsetX   int
	OffsetY   int
	ShapeSize int
}

func (b *SchemaBead) Render() template.HTML {
	var buf strings.Builder
	buf.Grow(100)

	buf.WriteString(`<rect width="`)
	buf.WriteString(strconv.Itoa(b.ShapeSize))
	buf.WriteString(`" height="`)
	buf.WriteString(strconv.Itoa(b.ShapeSize))
	buf.WriteString(`" rx="`)
	buf.WriteString(strconv.Itoa(b.ShapeSize))
	buf.WriteString(`" ry="`)
	buf.WriteString(strconv.Itoa(b.ShapeSize))
	buf.WriteString(`" x="`)
	buf.WriteString(strconv.Itoa(b.OffsetX))
	buf.WriteString(`" y="`)
	buf.WriteString(strconv.Itoa(b.OffsetY))
	buf.WriteString(`" fill="`)
	buf.WriteString(b.Color)
	buf.WriteString(`"/>`)

	return template.HTML(buf.String())
}

func beadsGrid(data *templates.SchemaPreviewData, fromX, toX, fromY, toY int) chan *SchemaBead {
	totalItems := (toX - fromX + 1) * (toY - fromY + 1)
	bufferSize := min(totalItems/20, 100)
	ch := make(chan *SchemaBead, bufferSize)

	initialOffsetX := -data.SizeLeft * int(data.BeadSize)
	initialOffsetY := -data.SizeTop * int(data.BeadSize)

	go func() {
		defer close(ch)

		var builder strings.Builder
		builder.Grow(20)

		for x := fromX; x <= toX; x++ {
			for y := fromY; y <= toY; y++ {
				builder.Reset()
				builder.WriteString(strconv.Itoa(x))
				builder.WriteByte(':')
				builder.WriteString(strconv.Itoa(y))
				coord := builder.String()

				color, ok := data.Beads[coord]

				if !ok {
					continue
				}

				ch <- &SchemaBead{
					OffsetX:   initialOffsetX + (x * int(data.BeadSize)),
					OffsetY:   initialOffsetY + (y * int(data.BeadSize)),
					Color:     color,
					ShapeSize: int(data.ShapeSize),
				}
			}
		}
	}()

	return ch
}
