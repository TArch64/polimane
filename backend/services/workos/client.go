package workos

import (
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/services/jwk"
)

type Client struct {
	UserManagement *usermanagement.Client
	MFA            *mfa.Client
	env            *env.Environment
	jwk            jwk.Interface
}

func Provider(environment *env.Environment, jwk jwk.Interface) *Client {
	usermanagement.SetAPIKey(environment.WorkOS.ApiKey)
	mfa.SetAPIKey(environment.WorkOS.ApiKey)

	return &Client{
		UserManagement: usermanagement.DefaultClient,
		MFA:            mfa.DefaultClient,
		env:            environment,
		jwk:            jwk,
	}
}
