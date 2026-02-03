package model

type AccessLevel uint8

const (
	AccessNone AccessLevel = iota
	AccessRead
	AccessWrite
	AccessAdmin
)
