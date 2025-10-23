package model

type User struct {
	*Identifiable
	WorkosID string `json:"-"`

	// Relations
	Schemas []*Schema `gorm:"many2many:user_schemas" json:"-"`
}
