package model

const TypeUser = "USER"

type User struct {
	*Base
	Username     string `dynamo:"sk,range" json:"username"`
	PasswordHash string `dynamo:"password_hash," json:"-"`
}
