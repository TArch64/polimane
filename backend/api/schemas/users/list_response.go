package users

import (
	"polimane/backend/model"
)

type ListResponse struct {
	Users       []*ListUser       `json:"users"`
	Invitations []*ListInvitation `json:"invitations"`
}

type ListAccessable struct {
	Access       model.AccessLevel `json:"access"`
	IsEvenAccess bool              `json:"isEvenAccess"`
	IsAllAccess  bool              `json:"isAllAccess"`
}

func (l *ListAccessable) AfterScan() error {
	if !l.IsEvenAccess || !l.IsAllAccess {
		l.Access = model.AccessNone
	}
	return nil
}

type ListUser struct {
	*ListAccessable
	ID        model.ID `json:"id"`
	Email     string   `json:"email"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
}

func NewListUser(user *model.User, access model.AccessLevel) *ListUser {
	return &ListUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,

		ListAccessable: &ListAccessable{
			Access:       access,
			IsEvenAccess: true,
			IsAllAccess:  true,
		},
	}
}

type ListInvitation struct {
	*ListAccessable
	Email string `json:"email"`
}

func NewListInvitation(email string, access model.AccessLevel) *ListInvitation {
	return &ListInvitation{
		Email: email,

		ListAccessable: &ListAccessable{
			Access:       access,
			IsEvenAccess: true,
			IsAllAccess:  true,
		},
	}
}
