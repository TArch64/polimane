package users

import (
	"polimane/backend/model"
)

type listResponse struct {
	Users       []*listUser       `json:"users"`
	Invitations []*listInvitation `json:"invitations"`
}

type listUser struct {
	ID             model.ID          `json:"id"`
	Email          string            `json:"email"`
	FirstName      string            `json:"firstName"`
	LastName       string            `json:"lastName"`
	Access         model.AccessLevel `json:"access"`
	IsUnevenAccess bool              `json:"isUnevenAccess"`
}

func newUserListItem(user *model.User, access model.AccessLevel) *listUser {
	return &listUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Access:    access,
	}
}

type listInvitation struct {
	Email          string            `json:"email"`
	Access         model.AccessLevel `json:"access"`
	IsUnevenAccess bool              `json:"isUnevenAccess"`
}
