package model

const (
	BeadCircle = "circle"
	BeadBugle  = "bugle"
	BeadRef    = "ref"
)

type SchemaBead struct {
	Kind   string            `validate:"required,oneof=circle bugle ref" json:"kind"`
	Circle *SchemaCircleBead `validate:"omitempty,dive" json:"circle,omitempty"`
	Bugle  *SchemaBugleBead  `validate:"omitempty,dive" json:"bugle,omitempty"`
	Ref    *SchemaRefBead    `validate:"omitempty,dive" json:"ref,omitempty"`
}

func (s *SchemaBead) GetColor() string {
	if s.Circle != nil {
		return s.Circle.Color
	}
	if s.Bugle != nil {
		return s.Bugle.Color
	}
	return ""
}

func (s *SchemaBead) GetSpan() *SchemaBeadSpan {
	if s.Bugle != nil {
		return s.Bugle.Span
	}
	return nil
}

type SchemaCircleBead struct {
	Color string `validate:"required,iscolor" json:"color"`
}

type SchemaBugleBead struct {
	Color string          `validate:"required,iscolor" json:"color"`
	Span  *SchemaBeadSpan `validate:"required,dive" json:"span"`
}

type SchemaRefBead struct {
	To string `validate:"required" json:"to"`
}

type SchemaBeadSpan struct {
	X uint8 `validate:"required,gte=0,lte=255" json:"x"`
	Y uint8 `validate:"required,gte=0,lte=255" json:"y"`
}
