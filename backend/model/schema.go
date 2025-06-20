package model

import "encoding/json"

const SKSchema = "SCHEMA"
const IndexSchemaID = "SchemaIdIndex"

type SchemaObject struct {
	ID string `json:"id" validate:"required,ulid"`
}

type SchemaPattern struct {
	*SchemaObject
	Name    string       `json:"name" validate:"required"`
	Type    string       `json:"type" validate:"required,oneof=square diamond"`
	Content []*SchemaRow `json:"content" validate:"required,dive,required"`
}

type SchemaRow struct {
	*SchemaObject
	Content []*SchemaBead `json:"content" validate:"required,dive,required"`
}

type SchemaBead struct {
	*SchemaObject
	Color string `json:"color" validate:"omitempty,iscolor"`
}

type SchemaContent []*SchemaPattern

type Schema struct {
	*Base
	Name    string        `dynamo:"Name"`
	Palette []string      `dynamo:"Palette"`
	Content SchemaContent `dynamo:"Content"`
}

func (u *Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID      string        `json:"id"`
		Name    string        `json:"name"`
		Palette []string      `json:"palette,omitempty"`
		Content SchemaContent `json:"content,omitempty"`
	}{
		ID:      u.SK.Value(),
		Name:    u.Name,
		Palette: u.Palette,
		Content: u.Content,
	})
}
