package model

type User struct {
	*Identifiable
	WorkosID string    `json:"-"`
	Schemas  []*Schema `gorm:"many2many:user_schemas" json:"-"`
}
