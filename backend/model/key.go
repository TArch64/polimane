package model

import "strings"

type Key string

const KeyDivider = "#"

func NewKey(keyType, keyValue string) Key {
	return Key(keyType + KeyDivider + keyValue)
}

func (k Key) String() string {
	return string(k)
}

func (k Key) Entries() []string {
	return strings.Split(string(k), KeyDivider)
}

func (k Key) Type() string {
	return k.Entries()[0]
}

func (k Key) Value() string {
	return k.Entries()[1]
}
