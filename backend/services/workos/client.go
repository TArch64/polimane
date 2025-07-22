package workos

import (
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/services/jwk"
)

type Client struct {
	UserManagement *usermanagement.Client
	env            *env.Environment
	jwk            jwk.Interface
}

func Provider(environment *env.Environment, jwk jwk.Interface) *Client {
	return &Client{
		UserManagement: usermanagement.NewClient(environment.WorkOS.ApiKey),
		env:            environment,
		jwk:            jwk,
	}
}
