package model

type User struct {
	*Identifiable
	*SoftDeletable
	WorkosID  string `json:"-"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	// Relations
	Schemas      []*Schema         `gorm:"many2many:user_schemas" json:"-"`
	Subscription *UserSubscription `gorm:"foreignKey:ID"`
}
