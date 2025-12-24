package model

import (
	"strconv"
	"strings"
)

type SchemaCoord string

func (c SchemaCoord) MustParseInt() (int, int) {
	parts := strings.SplitN(string(c), ":", 2)
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}
