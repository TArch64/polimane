package model

import "encoding/json"

const PKUser = "USER"
const SKUser = "USERNAME"
const IndexUserName = "UserNameIndex"

type User struct {
	*Base
	PasswordHash string `dynamo:"PasswordHash"`
}

type NewUserOptions struct {
	Username     string
	PasswordHash string
}

func NewUser(options *NewUserOptions) *User {
	return &User{
		Base: &Base{
			ID: NewID(PKUser),
			SK: NewKey(SKUser, options.Username),
		},
		PasswordHash: options.PasswordHash,
	}
}

func (u *User) Username() string {
	return u.SK.Value()
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}{
		ID:       u.ID.Value(),
		Username: u.Username(),
	})
}
