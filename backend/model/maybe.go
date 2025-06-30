package model

type Maybe[T any] struct {
	Value T
	Err   error
}

func NewMaybeValue[T any](value T) *Maybe[T] {
	return &Maybe[T]{Value: value}
}

func NewMaybeError[T any](err error) *Maybe[T] {
	return &Maybe[T]{Err: err}
}

func (m *Maybe[T]) IsOk() bool {
	return m.Err == nil
}
