package workos

import (
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"go.uber.org/fx"

	"polimane/backend/env"
)

type ClientOptions struct {
	fx.In
	Env *env.Environment
}

type Client struct {
	UserManagement *usermanagement.Client
	MFA            *mfa.Client
	env            *env.Environment
}

func Provider(options ClientOptions) *Client {
	usermanagement.SetAPIKey(options.Env.WorkOS.ApiKey)
	mfa.SetAPIKey(options.Env.WorkOS.ApiKey)

	return &Client{
		UserManagement: usermanagement.DefaultClient,
		MFA:            mfa.DefaultClient,
		env:            options.Env,
	}
}
