package model

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

type ID Key

func NewID(modelType string) ID {
	return ID(NewKey(modelType, strings.ToLower(ulid.Make().String())))
}

func (i ID) String() string {
	return string(i)
}

func (i ID) Key() Key {
	return Key(i)
}

func (i ID) Type() string {
	return Key(i).Type()
}

func (i ID) Value() string {
	return Key(i).Value()
}

func (i ID) ULID() ulid.ULID {
	return ulid.MustParse(i.Value())
}
