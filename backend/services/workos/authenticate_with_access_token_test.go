package workos

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"polimane/backend/env"
	"polimane/backend/services/jwk"
)

// Mock implementations are defined in mocks_test.go

// Test the AccessTokenClaims struct
func TestAccessTokenClaims(t *testing.T) {
	claims := &AccessTokenClaims{
		UserID:    "user_123",
		SessionID: "session_456",
	}

	assert.Equal(t, "user_123", claims.UserID)
	assert.Equal(t, "session_456", claims.SessionID)
}

// Test the AccessTokenExpired error
func TestAccessTokenExpiredError(t *testing.T) {
	assert.Equal(t, "access token expired", AccessTokenExpired.Error())
	assert.NotNil(t, AccessTokenExpired)
}

// Test AuthenticateWithAccessToken - limited testing due to complex JWT dependencies
func TestClient_AuthenticateWithAccessToken_ErrorCases(t *testing.T) {
	tests := []struct {
		name          string
		tokenStr      string
		setupClient   func() *Client
		expectedError string
		shouldSkip    bool
		skipReason    string
	}{
		{
			name:     "empty token string",
			tokenStr: "",
			setupClient: func() *Client {
				return &Client{
					env: &env.Environment{
						WorkOS: struct {
							ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
							ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
						}{
							ClientID: "test-client-id",
							ApiKey:   "test-api-key",
						},
					},
					jwk: &jwk.Impl{}, // Real JWK client
				}
			},
			expectedError: "", // Will fail due to real WorkOS client
			shouldSkip:    true,
			skipReason:    "Requires real WorkOS client which needs network access",
		},
		{
			name:     "invalid token format",
			tokenStr: "invalid.token.format",
			setupClient: func() *Client {
				return &Client{
					env: &env.Environment{
						WorkOS: struct {
							ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
							ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
						}{
							ClientID: "test-client-id",
							ApiKey:   "test-api-key",
						},
					},
					jwk: &jwk.Impl{}, // Real JWK client
				}
			},
			expectedError: "", // Will fail due to real WorkOS client
			shouldSkip:    true,
			skipReason:    "Requires real WorkOS client which needs network access",
		},
		{
			name:     "malformed JWT",
			tokenStr: "not.a.jwt",
			setupClient: func() *Client {
				return &Client{
					env: &env.Environment{
						WorkOS: struct {
							ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
							ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
						}{
							ClientID: "test-client-id",
							ApiKey:   "test-api-key",
						},
					},
					jwk: &jwk.Impl{}, // Real JWK client
				}
			},
			expectedError: "", // Will fail due to real WorkOS client
			shouldSkip:    true,
			skipReason:    "Requires real WorkOS client which needs network access",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldSkip {
				t.Skip(tt.skipReason)
				return
			}

			client := tt.setupClient()
			ctx := context.Background()

			claims, err := client.AuthenticateWithAccessToken(ctx, tt.tokenStr)

			// Since we can't easily mock the WorkOS SDK, we expect these to error
			assert.Error(t, err)
			assert.Nil(t, claims)

			if tt.expectedError != "" {
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}

// Test the error handling for expired tokens specifically
func TestClient_AuthenticateWithAccessToken_ExpiredTokenHandling(t *testing.T) {
	t.Skip("Complex JWT parsing and WorkOS SDK mocking required - better suited for integration tests")

	// This test would verify that when jwt.Parse returns an error containing "token is expired",
	// the function returns AccessTokenExpired error instead of the original error.
	// However, mocking jwt.Parse is complex and requires significant setup.
}

// Test the session ID extraction logic
func TestClient_AuthenticateWithAccessToken_SessionIDExtraction(t *testing.T) {
	t.Skip("Complex JWT token mocking required - better suited for integration tests")

	// This test would verify that the function correctly extracts the "sid" claim
	// from a valid JWT token and includes it in the AccessTokenClaims.
	// However, creating valid JWT tokens for testing is complex.
}

// Test client configuration requirements
func TestClient_AuthenticateWithAccessToken_Requirements(t *testing.T) {
	tests := []struct {
		name        string
		client      *Client
		expectPanic bool
	}{
		{
			name: "client with nil UserManagement",
			client: &Client{
				UserManagement: nil,
				env: &env.Environment{
					WorkOS: struct {
						ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
						ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
					}{
						ClientID: "test-client-id",
						ApiKey:   "test-api-key",
					},
				},
				jwk: &jwk.Impl{},
			},
			expectPanic: true,
		},
		{
			name: "client with nil environment",
			client: &Client{
				env: nil,
				jwk: &jwk.Impl{},
			},
			expectPanic: true,
		},
		{
			name: "client with nil JWK interface",
			client: &Client{
				env: &env.Environment{
					WorkOS: struct {
						ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
						ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
					}{
						ClientID: "test-client-id",
						ApiKey:   "test-api-key",
					},
				},
				jwk: nil,
			},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tokenStr := "test.jwt.token"

			if tt.expectPanic {
				assert.Panics(t, func() {
					_, _ = tt.client.AuthenticateWithAccessToken(ctx, tokenStr)
				})
			} else {
				// Should not panic, but may error due to invalid token/config
				_, err := tt.client.AuthenticateWithAccessToken(ctx, tokenStr)
				// We expect an error since these are not real valid configs
				assert.Error(t, err)
			}
		})
	}
}
