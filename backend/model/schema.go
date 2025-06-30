package model

import "encoding/json"

const PKSchemaPrefix = "SCHEMA"
const SKSchema = "#SCHEMA"

type SchemaContent []interface{}

type Schema struct {
	*Base
	UserIDs []PrimaryKey  `json:"userIds" dynamo:"UserIDs,set"`
	Name    string        `json:"name" dynamo:"Name"`
	Palette []string      `json:"palette" dynamo:"Palette"`
	Content SchemaContent `json:"content" dynamo:"Content"`
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID      string        `json:"id"`
		Name    string        `json:"name"`
		Palette []string      `json:"palette,omitempty"`
		Content SchemaContent `json:"content,omitempty"`
	}{
		ID:      s.PK.Value(),
		Name:    s.Name,
		Palette: s.Palette,
		Content: s.Content,
	})
}
