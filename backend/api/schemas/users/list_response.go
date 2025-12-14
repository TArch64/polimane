package users

import (
	"polimane/backend/model"
)

type ListResponse struct {
	Users       []*ListUser       `json:"users"`
	Invitations []*ListInvitation `json:"invitations"`
}

type ListUser struct {
	ID             model.ID          `json:"id"`
	Email          string            `json:"email"`
	FirstName      string            `json:"firstName"`
	LastName       string            `json:"lastName"`
	Access         model.AccessLevel `json:"access"`
	IsUnevenAccess bool              `json:"isUnevenAccess"`
}

func NewUserListItem(user *model.User, access model.AccessLevel) *ListUser {
	return &ListUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Access:    access,
	}
}

type ListInvitation struct {
	Email          string            `json:"email"`
	Access         model.AccessLevel `json:"access"`
	IsUnevenAccess bool              `json:"isUnevenAccess"`
}
