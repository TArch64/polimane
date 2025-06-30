package model

import (
	"encoding/json"
	"slices"

	"github.com/guregu/dynamo/v2"
)

const PKUserPrefix = "USER"
const SKUserPrefix = "USERNAME"
const IndexUserName = "UserNameIndex"

type User struct {
	*Base
	SchemaIDs    []PrimaryKey `dynamo:"SchemaIDs,set"`
	PasswordHash string       `dynamo:"PasswordHash"`
}

type NewUserOptions struct {
	Username     string
	PasswordHash string
}

func NewUser(options *NewUserOptions) *User {
	return &User{
		Base: &Base{
			PK: RandomID(PKUserPrefix),
			SK: NewKey(SKUserPrefix, options.Username),
		},
		PasswordHash: options.PasswordHash,
	}
}

func (u *User) Username() string {
	return u.SK.Value()
}

func (u *User) CheckSchemaAccess(checkingId ID) error {
	for _, id := range u.SchemaIDs {
		if id.PK() == checkingId {
			return nil
		}
	}

	return dynamo.ErrNotFound
}

func (u *User) AddSchemaID(key PrimaryKey) {
	u.SchemaIDs = append(u.SchemaIDs, key)
}

func (u *User) DeleteSchemaID(removingId ID) {
	u.SchemaIDs = slices.DeleteFunc(u.SchemaIDs, func(id PrimaryKey) bool {
		return id.PK() == removingId
	})
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}{
		ID:       u.PK.Value(),
		Username: u.Username(),
	})
}
