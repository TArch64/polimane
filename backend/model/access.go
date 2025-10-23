package model

type AccessLevel uint8

const (
	AccessRead  AccessLevel = 1
	AccessWrite AccessLevel = 2
	AccessAdmin AccessLevel = 3
)
