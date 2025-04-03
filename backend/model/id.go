package model

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID(modelType string) string {
	return modelType + "#" + ulid.Make().String()
}

func (id ID) ULID() ulid.ULID {
	ulidStr := strings.Split(string(id), "#")[1]
	return ulid.MustParse(ulidStr)
}
