package model

func ptr[V any](v V) *V {
	return &v
}
