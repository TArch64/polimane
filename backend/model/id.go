package model

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

type ID Key

func NewID(modelType string) ID {
	return ID(NewKey(modelType, strings.ToLower(ulid.Make().String())))
}

func (id ID) Type() string {
	return Key(id).Type()
}

func (id ID) Value() string {
	return Key(id).Value()
}

func (id ID) ULID() ulid.ULID {
	return ulid.MustParse(id.Value())
}
