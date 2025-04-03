package model

import "github.com/oklog/ulid/v2"

type Base struct {
	ID ulid.ULID `dynamo:"pk,hash" json:"id"`
}

func TypeFilter(modelType string) (string, interface{}) {
	return "begins_with(pk, ?)", modelType
}
