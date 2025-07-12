package model

import (
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model/modelbase"
)

type User struct {
	*modelbase.Identifiable
	WorkosID string    `gorm:"not null;uniqueIndex;size:32" json:"-"`
	Schemas  []*Schema `gorm:"many2many:user_schemas;constraint:OnDelete:Cascade;" json:"-"`
}

func (u *User) AsFullUser(workosUser *usermanagement.User) *FullUser {
	return &FullUser{
		ID:            u.ID,
		Email:         workosUser.Email,
		FirstName:     workosUser.FirstName,
		LastName:      workosUser.LastName,
		EmailVerified: workosUser.EmailVerified,
	}
}

type FullUser struct {
	ID            modelbase.ID `json:"id"`
	Email         string       `json:"email"`
	FirstName     string       `json:"firstName"`
	LastName      string       `json:"lastName"`
	EmailVerified bool         `json:"emailVerified"`
}
