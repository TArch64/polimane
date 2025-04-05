package model

import "encoding/json"

const SKSchema = "SCHEMA"
const IndexSchemaID = "SchemaIdIndex"

type SchemaContent map[string]struct{}

type Schema struct {
	*Base
	Content SchemaContent
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
