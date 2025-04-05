package model

type Base struct {
	ID ID  `dynamo:"PK,hash"`
	SK Key `dynamo:"SK,range"`
}
