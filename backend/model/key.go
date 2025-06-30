package model

import (
	"strings"
)

type Key string

const KeyDivider = "#"

func NewKey(keyType, keyValue string) Key {
	return Key(keyType + KeyDivider + keyValue)
}

func StringToKey(str string) Key {
	parts := strings.Split(str, KeyDivider)
	if len(parts) != 2 {
		// almost impossible to reach this, but just in case
		panic("Invalid key format, expected 'type#value'")
	}

	return Key(str)
}

func (k Key) String() string {
	return string(k)
}

func (k Key) Entries() [2]string {
	parts := strings.Split(string(k), KeyDivider)
	return [2]string{parts[0], parts[1]}
}

func (k Key) Type() string {
	return k.Entries()[0]
}

func (k Key) Value() string {
	return k.Entries()[1]
}
