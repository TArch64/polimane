package auth

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
)

func TestApiLogin(t *testing.T) {
	t.Run("returns authorization URL successfully", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedURL, _ := url.Parse("https://api.workos.com/auth")

		controller := &Controller{
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
				AppProtocol: "https",
				AppDomain:   "api.test.com",
			},
		}

		mockUserManagement.On("GetAuthorizationURL", mock.AnythingOfType("usermanagement.GetAuthorizationURLOpts")).Return(expectedURL, nil)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/login", controller.apiLogin)
		req := httptest.NewRequest("GET", "/login", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Check response body
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedURL.String(), response["url"])

		mockUserManagement.AssertExpectations(t)
	})

	t.Run("handles workos client error", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedError := assert.AnError

		controller := &Controller{
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
				AppProtocol: "https",
				AppDomain:   "api.test.com",
			},
		}

		mockUserManagement.On("GetAuthorizationURL", mock.Anything).Return((*url.URL)(nil), expectedError)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/login", controller.apiLogin)
		req := httptest.NewRequest("GET", "/login", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
	})

	t.Run("uses correct redirect URI", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedURL, _ := url.Parse("https://api.workos.com/auth")

		controller := &Controller{
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "client-123",
				},
				AppProtocol: "http",
				AppDomain:   "localhost:8080",
			},
		}

		mockUserManagement.On("GetAuthorizationURL", mock.MatchedBy(func(opts usermanagement.GetAuthorizationURLOpts) bool {
			return opts.RedirectURI == "http://api.localhost:8080/api/auth/login/complete"
		})).Return(expectedURL, nil)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/login", controller.apiLogin)
		req := httptest.NewRequest("GET", "/login", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
	})

	t.Run("uses authkit provider", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedURL, _ := url.Parse("https://api.workos.com/auth")

		controller := &Controller{
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "client-123",
				},
				AppProtocol: "https",
				AppDomain:   "example.com",
			},
		}

		mockUserManagement.On("GetAuthorizationURL", mock.MatchedBy(func(opts usermanagement.GetAuthorizationURLOpts) bool {
			return opts.Provider == "authkit"
		})).Return(expectedURL, nil)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/login", controller.apiLogin)
		req := httptest.NewRequest("GET", "/login", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
	})

	t.Run("returns JSON response with correct structure", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedURL, _ := url.Parse("https://test.workos.com/auth?param=value")

		controller := &Controller{
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client",
				},
				AppProtocol: "https",
				AppDomain:   "api.example.com",
			},
		}

		mockUserManagement.On("GetAuthorizationURL", mock.Anything).Return(expectedURL, nil)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/login", controller.apiLogin)
		req := httptest.NewRequest("GET", "/login", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify Content-Type header
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		// Check response structure
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Len(t, response, 1)
		assert.Contains(t, response, "url")
		assert.Equal(t, expectedURL.String(), response["url"])

		mockUserManagement.AssertExpectations(t)
	})
}
