package model

type Get[T, V any] func(target T) V
type Set[T, V any] func(target T, value V)

type Accessor[T, V any] struct {
	Get Get[T, V]
	Set Set[T, V]
}

func NewAccessor[T, V any](get Get[T, V], set Set[T, V]) *Accessor[T, V] {
	return &Accessor[T, V]{
		Get: get,
		Set: set,
	}
}
