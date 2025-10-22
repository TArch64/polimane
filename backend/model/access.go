package model

type AccessLevel string

const (
	AccessRead  AccessLevel = "read"
	AccessWrite AccessLevel = "write"
	AccessAdmin AccessLevel = "admin"
)
