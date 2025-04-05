package model

import "encoding/json"

const SKSchema = "SCHEMA"
const IndexSchemaID = "SchemaIdIndex"

type SchemaContent map[string]interface{}

type Schema struct {
	*Base
	Name    string        `dynamo:"Name"`
	Content SchemaContent `dynamo:"Content"`
}

func (u *Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID      string        `json:"id"`
		Content SchemaContent `json:"content"`
	}{
		ID:      u.SK.Value(),
		Content: u.Content,
	})
}
