package auth

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	"polimane/backend/services/workos"
)

func TestApiLoginComplete(t *testing.T) {
	t.Run("completes login successfully", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockUsers := &MockUsersClient{}

		testUserID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &modelbase.Identifiable{ID: testUserID},
		}

		authResponse := usermanagement.AuthenticateResponse{
			User: usermanagement.User{
				ID: "workos-user-123",
			},
			AccessToken:  "access-token-123",
			RefreshToken: "refresh-token-456",
		}

		updatedUser := usermanagement.User{
			ID:         "workos-user-123",
			ExternalID: testUserID.String(),
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
				AppProtocol: "https",
				AppDomain:   "app.test.com",
			},
			users: mockUsers,
		}

		mockUserManagement.On("AuthenticateWithCode", mock.Anything, mock.MatchedBy(func(opts usermanagement.AuthenticateWithCodeOpts) bool {
			return opts.ClientID == "test-client-id" && opts.Code == "test-auth-code"
		})).Return(authResponse, nil)

		mockUsers.On("CreateIfNeeded", mock.Anything, "workos-user-123").Return(testUser, nil)

		mockUserManagement.On("UpdateUser", mock.Anything, mock.MatchedBy(func(opts usermanagement.UpdateUserOpts) bool {
			return opts.User == "workos-user-123" && opts.ExternalID == testUserID.String()
		})).Return(updatedUser, nil)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/complete", controller.apiLoginComplete)
		req := httptest.NewRequest("GET", "/complete?code=test-auth-code", nil)
		req.Header.Set("User-Agent", "test-browser/1.0")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 302, resp.StatusCode)

		// Check redirect URL
		location := resp.Header.Get("Location")
		assert.NotEmpty(t, location)

		parsedURL, err := url.Parse(location)
		assert.NoError(t, err)
		assert.Equal(t, "https", parsedURL.Scheme)
		assert.Equal(t, "app.test.com", parsedURL.Host)
		assert.Equal(t, "/auth/complete", parsedURL.Path)

		// Check query parameters
		query := parsedURL.Query()
		assert.Equal(t, "access-token-123", query.Get("access-token"))
		assert.Equal(t, "refresh-token-456", query.Get("refresh-token"))

		mockUserManagement.AssertExpectations(t)
		mockUsers.AssertExpectations(t)
	})

	//t.Run("handles missing code parameter", func(t *testing.T) {
	//	// Arrange
	//	controller := &Controller{
	//		workosClient: &workos.Client{
	//			UserManagement: &MockUserManagement{},
	//		},
	//		env: &env.Environment{
	//			WorkOS: struct {
	//				ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
	//				ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
	//			}{},
	//		},
	//	}
	//
	//	// Create fiber app and request without code parameter
	//	app := fiber.New()
	//	app.Get("/complete", controller.apiLoginComplete)
	//	req := httptest.NewRequest("GET", "/complete", nil)
	//
	//	// Act
	//	resp, err := app.Test(req)
	//
	//	// Assert
	//	assert.NoError(t, err)
	//	assert.Equal(t, 400, resp.StatusCode)
	//})

	t.Run("handles WorkOS authentication error", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedError := assert.AnError

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
			},
		}

		var emptyResponse usermanagement.AuthenticateResponse
		mockUserManagement.On("AuthenticateWithCode", mock.Anything, mock.Anything).Return(emptyResponse, expectedError)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/complete", controller.apiLoginComplete)
		req := httptest.NewRequest("GET", "/complete?code=invalid-code", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
	})

	t.Run("handles user creation error", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockUsers := &MockUsersClient{}
		expectedError := assert.AnError

		authResponse := usermanagement.AuthenticateResponse{
			User: usermanagement.User{
				ID: "workos-user-123",
			},
			AccessToken:  "access-token-123",
			RefreshToken: "refresh-token-456",
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
			},
			users: mockUsers,
		}

		mockUserManagement.On("AuthenticateWithCode", mock.Anything, mock.Anything).Return(authResponse, nil)
		mockUsers.On("CreateIfNeeded", mock.Anything, "workos-user-123").Return((*model.User)(nil), expectedError)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/complete", controller.apiLoginComplete)
		req := httptest.NewRequest("GET", "/complete?code=test-code", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
		mockUsers.AssertExpectations(t)
	})

	t.Run("handles user update error", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockUsers := &MockUsersClient{}
		expectedError := assert.AnError

		testUserID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &modelbase.Identifiable{ID: testUserID},
		}

		authResponse := usermanagement.AuthenticateResponse{
			User: usermanagement.User{
				ID: "workos-user-123",
			},
			AccessToken:  "access-token-123",
			RefreshToken: "refresh-token-456",
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
			},
			users: mockUsers,
		}

		mockUserManagement.On("AuthenticateWithCode", mock.Anything, mock.Anything).Return(authResponse, nil)
		mockUsers.On("CreateIfNeeded", mock.Anything, "workos-user-123").Return(testUser, nil)

		var emptyUser usermanagement.User
		mockUserManagement.On("UpdateUser", mock.Anything, mock.Anything).Return(emptyUser, expectedError)

		// Create fiber app and request
		app := fiber.New()
		app.Get("/complete", controller.apiLoginComplete)
		req := httptest.NewRequest("GET", "/complete?code=test-code", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
		mockUsers.AssertExpectations(t)
	})

	t.Run("forwards User-Agent header to WorkOS", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockUsers := &MockUsersClient{}

		testUserID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &modelbase.Identifiable{ID: testUserID},
		}

		authResponse := usermanagement.AuthenticateResponse{
			User: usermanagement.User{
				ID: "workos-user-123",
			},
			AccessToken:  "access-token-123",
			RefreshToken: "refresh-token-456",
		}

		updatedUser := usermanagement.User{
			ID:         "workos-user-123",
			ExternalID: testUserID.String(),
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
				},
				AppProtocol: "https",
				AppDomain:   "app.test.com",
			},
			users: mockUsers,
		}

		mockUserManagement.On("AuthenticateWithCode", mock.Anything, mock.MatchedBy(func(opts usermanagement.AuthenticateWithCodeOpts) bool {
			return opts.UserAgent == "custom-browser/2.0"
		})).Return(authResponse, nil)

		mockUsers.On("CreateIfNeeded", mock.Anything, "workos-user-123").Return(testUser, nil)
		mockUserManagement.On("UpdateUser", mock.Anything, mock.Anything).Return(updatedUser, nil)

		// Create fiber app and request with custom User-Agent
		app := fiber.New()
		app.Get("/complete", controller.apiLoginComplete)
		req := httptest.NewRequest("GET", "/complete?code=test-code", nil)
		req.Header.Set("User-Agent", "custom-browser/2.0")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 302, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
		mockUsers.AssertExpectations(t)
	})
}

func TestLoginCompleteQuery(t *testing.T) {
	t.Run("struct initialization", func(t *testing.T) {
		query := loginCompleteQuery{
			Code: "test-code-123",
		}

		assert.Equal(t, "test-code-123", query.Code)
	})

	t.Run("zero value", func(t *testing.T) {
		var query loginCompleteQuery

		assert.Equal(t, "", query.Code)
	})
}
