package model

type Get[T, V any] func(target T) V
type Set[T, V any] func(target T, value V)
