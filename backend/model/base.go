package model

type Base struct {
	ID ID  `dynamo:"PK,hash"`
	SK Key `dynamo:"SK,range"`
}

func TypeFilter(modelType string) (string, interface{}) {
	return "begins_with(PK, ?)", modelType
}
