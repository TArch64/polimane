package model

type Base struct {
	ID ID  `dynamo:"PK,hash"`
	SK Key `dynamo:"SK,range"`
}

const IfPKExists = "attribute_exists(PK)"
const IfSKExists = "attribute_exists(SK)"
const IfKeyExists = IfPKExists + " and " + IfSKExists
