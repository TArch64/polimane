package model

type Base struct {
	ID ID `dynamo:"pk,hash" json:"id"`
}

func TypeFilter(modelType string) (string, interface{}) {
	return "begins_with(pk, ?)", modelType
}
