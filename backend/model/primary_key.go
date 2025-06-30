package model

import (
	"strings"

	"github.com/guregu/dynamo/v2"
)

type PrimaryKey string

const PrimaryKeyDivider = "&"

func NewPrimaryKey(pk, sk ID) PrimaryKey {
	return PrimaryKey(pk + PrimaryKeyDivider + sk)
}

func (k PrimaryKey) String() string {
	return string(k)
}

func (k PrimaryKey) Entries() [2]ID {
	parts := strings.Split(k.String(), PrimaryKeyDivider)
	return [2]ID{ID(parts[0]), ID(parts[1])}
}

func (k PrimaryKey) PK() ID {
	return k.Entries()[0]
}

func (k PrimaryKey) SK() ID {
	return k.Entries()[1]
}

func (k PrimaryKey) Keys() dynamo.Keyed {
	entries := k.Entries()
	return dynamo.Keys{entries[0], entries[1]}
}
