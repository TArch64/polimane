package model

import "encoding/json"

const PKUser = "USER"
const SKUser = "USERNAME"

type User struct {
	*Base
	PasswordHash string `dynamo:"PasswordHash"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}{
		ID:       u.ID.Value(),
		Username: u.SK.Value(),
	})
}
