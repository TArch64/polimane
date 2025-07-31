package workos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/services/jwk"
)

func TestProvider(t *testing.T) {
	tests := []struct {
		name         string
		environment  *env.Environment
		jwkInterface jwk.Interface
		expectNil    bool
	}{
		{
			name: "successful client creation",
			environment: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
					ApiKey:   "test-api-key",
				},
			},
			jwkInterface: nil, // Using nil for simplicity
			expectNil:    false,
		},
		{
			name: "client creation with empty credentials",
			environment: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "",
					ApiKey:   "",
				},
			},
			jwkInterface: nil,
			expectNil:    false, // Provider doesn't validate credentials, just creates client
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := Provider(tt.environment, tt.jwkInterface)

			if tt.expectNil {
				assert.Nil(t, client)
			} else {
				assert.NotNil(t, client)
				assert.NotNil(t, client.UserManagement)
				assert.NotNil(t, client.MFA)
				assert.Equal(t, tt.environment, client.env)
				assert.Equal(t, tt.jwkInterface, client.jwk)

				// Verify that the clients are the default WorkOS clients
				assert.Equal(t, usermanagement.DefaultClient, client.UserManagement)
				assert.Equal(t, mfa.DefaultClient, client.MFA)
			}
		})
	}
}

func TestClient_Initialization(t *testing.T) {
	// Test that the Client struct has the expected fields
	environment := &env.Environment{
		WorkOS: struct {
			ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
			ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
		}{
			ClientID: "test-client-id",
			ApiKey:   "test-api-key",
		},
	}

	jwkInterface := (*jwk.Client)(nil) // Using nil jwk.Client pointer

	client := &Client{
		UserManagement: usermanagement.DefaultClient,
		MFA:            mfa.DefaultClient,
		env:            environment,
		jwk:            jwkInterface,
	}

	assert.NotNil(t, client.UserManagement)
	assert.NotNil(t, client.MFA)
	assert.Equal(t, environment, client.env)
	assert.Equal(t, jwkInterface, client.jwk)
}

func TestProvider_APIKeyConfiguration(t *testing.T) {
	// Test that Provider function sets API keys correctly
	environment := &env.Environment{
		WorkOS: struct {
			ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
			ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
		}{
			ClientID: "test-client-id",
			ApiKey:   "test-api-key-12345",
		},
	}

	// Note: We can't easily test that SetAPIKey was called with the correct values
	// since it's a side effect that modifies global state in the WorkOS library.
	// This test mainly verifies that Provider doesn't panic and returns a valid client.
	client := Provider(environment, nil)

	assert.NotNil(t, client)
	assert.Equal(t, environment, client.env)
	assert.Nil(t, client.jwk)

	// Verify client has the expected structure
	assert.NotNil(t, client.UserManagement)
	assert.NotNil(t, client.MFA)
}

func TestClient_FieldTypes(t *testing.T) {
	// Test that Client struct fields have the expected types
	client := &Client{}

	// Test that we can assign the expected types
	client.UserManagement = usermanagement.DefaultClient
	client.MFA = mfa.DefaultClient
	client.env = &env.Environment{}
	client.jwk = nil

	// Verify assignments worked
	assert.Equal(t, usermanagement.DefaultClient, client.UserManagement)
	assert.Equal(t, mfa.DefaultClient, client.MFA)
	assert.NotNil(t, client.env)
	assert.Nil(t, client.jwk)
}
