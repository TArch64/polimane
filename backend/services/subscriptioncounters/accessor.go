package subscriptioncounters

type accessor[T any] struct {
	Get func(target *T) uint16
	Set func(target *T, value uint16)
}
