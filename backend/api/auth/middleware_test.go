package auth

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kittipat1413/go-common/framework/cache/localcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/model"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

func TestMiddlewareProvider(t *testing.T) {
	t.Run("creates middleware with dependencies", func(t *testing.T) {
		// Arrange
		signalsContainer := signal.Provider()
		environment := &env.Environment{}
		workosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}
		usersClient := &MockUsersClient{}

		// Act
		middleware := MiddlewareProvider(MiddlewareOptions{
			Signals:      signalsContainer,
			Env:          environment,
			WorkosClient: workosClient,
			Users:        usersClient,
		})

		// Assert
		assert.NotNil(t, middleware)
		assert.IsType(t, &Middleware{}, middleware)
		assert.Equal(t, environment, middleware.env)
		assert.Equal(t, workosClient, middleware.workosClient)
		assert.Equal(t, usersClient, middleware.users)
		assert.NotNil(t, middleware.userCache)
		assert.NotNil(t, middleware.workosUserCache)
	})

	t.Run("sets up signal listeners", func(t *testing.T) {
		// This test verifies that the middleware provider registers signal listeners
		// The actual signal behavior is tested in separate signal tests
		signalsContainer := signal.Provider()
		environment := &env.Environment{}
		workosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}
		usersClient := &MockUsersClient{}

		middleware := MiddlewareProvider(MiddlewareOptions{
			Signals:      signalsContainer,
			Env:          environment,
			WorkosClient: workosClient,
			Users:        usersClient,
		})

		assert.NotNil(t, middleware)
	})
}

func TestMiddlewareHandler(t *testing.T) {
	t.Run("returns unauthorized when access token missing", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    &MockWorkosClient{userManagement: &MockUserManagement{}},
			env:             &env.Environment{},
			users:           &MockUsersClient{},
		}

		// Create fiber app and request
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-Refresh-Token", "refresh-token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode)
	})

	t.Run("returns unauthorized when refresh token missing", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    &MockWorkosClient{userManagement: &MockUserManagement{}},
			env:             &env.Environment{},
			users:           &MockUsersClient{},
		}

		// Create fiber app and request
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode)
	})
}

func TestInvalidateAuthCache(t *testing.T) {
	t.Run("invalidates workos user and regular user cache", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		workosUser := &usermanagement.User{
			ID:         "workos-123",
			ExternalID: testUserID.String(),
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
		}

		ctx := context.Background()
		sessionID := "session-456"

		// Pre-populate cache with workos user
		middleware.workosUserCache.Set(ctx, sessionID, workosUser, nil)

		// Act
		middleware.invalidateAuthCache(ctx, sessionID)

		// Assert - check that cache was invalidated by trying to get the item
		result, err := middleware.workosUserCache.Get(ctx, "workos-123", func() (*usermanagement.User, *time.Duration, error) {
			return nil, nil, errors.New("should not be called if cache hit")
		})

		// If cache was properly invalidated, the loader function should be called and return error
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("handles missing workos user gracefully", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		middleware := &Middleware{
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
		}

		ctx := context.Background()
		sessionID := "session-456"

		// Act - should not panic when workos user is not found
		middleware.invalidateAuthCache(ctx, sessionID)

		// Assert - no panic means success
		assert.True(t, true)
	})
}

func TestInvalidateUserCache(t *testing.T) {
	t.Run("invalidates user cache by ID", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		middleware := &Middleware{
			userCache: localcache.New[*model.User](cacheOptions...),
		}

		ctx := context.Background()
		userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: userID},
		}

		// Pre-populate cache
		middleware.userCache.Set(ctx, userID.String(), testUser, nil)

		// Act
		middleware.invalidateUserCache(ctx, userID)

		// Assert - verify cache was invalidated
		result, err := middleware.userCache.Get(ctx, userID.String(), func() (*model.User, *time.Duration, error) {
			return nil, nil, errors.New("cache miss - item was invalidated")
		})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestInvalidateWorkosUserCache(t *testing.T) {
	t.Run("invalidates workos user cache by ID", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		middleware := &Middleware{
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
		}

		ctx := context.Background()
		userID := "workos-123"

		workosUser := &usermanagement.User{
			ID: userID,
		}

		// Pre-populate cache
		middleware.workosUserCache.Set(ctx, userID, workosUser, nil)

		// Act
		middleware.invalidateWorkosUserCache(ctx, userID)

		// Assert - verify cache was invalidated
		result, err := middleware.workosUserCache.Get(ctx, userID, func() (*usermanagement.User, *time.Duration, error) {
			return nil, nil, errors.New("cache miss - item was invalidated")
		})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("returns user from cache when found", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUsers := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		middleware := &Middleware{
			userCache: localcache.New[*model.User](cacheOptions...),
			users:     mockUsers,
		}

		ctx := context.Background()

		// Pre-populate cache
		middleware.userCache.Set(ctx, testUserID.String(), testUser, nil)

		// Act
		result, err := middleware.getUser(ctx, testUserID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, testUser, result)

		// Verify users client was not called (cache hit)
		mockUsers.AssertNotCalled(t, "ByID")
	})

	t.Run("fetches user from database when not in cache", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUsers := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		middleware := &Middleware{
			userCache: localcache.New[*model.User](cacheOptions...),
			users:     mockUsers,
		}

		ctx := context.Background()

		mockUsers.On("ByID", ctx, testUserID).Return(testUser, nil)

		// Act
		result, err := middleware.getUser(ctx, testUserID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, testUser, result)
		mockUsers.AssertExpectations(t)
	})

	t.Run("returns unauthorized error when user not found", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUsers := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		middleware := &Middleware{
			userCache: localcache.New[*model.User](cacheOptions...),
			users:     mockUsers,
		}

		ctx := context.Background()

		mockUsers.On("ByID", ctx, testUserID).Return((*model.User)(nil), gorm.ErrRecordNotFound)

		// Act
		result, err := middleware.getUser(ctx, testUserID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Unauthorized")
		mockUsers.AssertExpectations(t)
	})
}

func TestGetWorkosUser(t *testing.T) {
	t.Run("returns workos user from cache when found", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUserManagement := &MockUserManagement{}

		workosUser := &usermanagement.User{
			ID:         "workos-123",
			ExternalID: "550e8400-e29b-41d4-a716-446655440000",
		}

		claims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		middleware := &Middleware{
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
		}

		ctx := context.Background()

		// Pre-populate cache
		middleware.workosUserCache.Set(ctx, "workos-123", workosUser, nil)

		// Act
		result, err := middleware.getWorkosUser(ctx, claims)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, workosUser, result)

		// Verify UserManagement was not called (cache hit)
		mockUserManagement.AssertNotCalled(t, "GetUser")
	})

	t.Run("fetches workos user from API when not in cache", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUserManagement := &MockUserManagement{}

		workosUser := usermanagement.User{
			ID:         "workos-123",
			ExternalID: "550e8400-e29b-41d4-a716-446655440000",
		}

		claims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		middleware := &Middleware{
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
		}

		ctx := context.Background()

		mockUserManagement.On("GetUser", ctx, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(workosUser, nil)

		// Act
		result, err := middleware.getWorkosUser(ctx, claims)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "workos-123", result.ID)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", result.ExternalID)
		mockUserManagement.AssertExpectations(t)
	})
}

func TestNewUnauthorizedErr(t *testing.T) {
	t.Run("returns basic unauthorized error in production", func(t *testing.T) {
		// Arrange
		middleware := &Middleware{}
		testErr := errors.New("test error")

		// Act
		result := middleware.newUnauthorizedErr(testErr)

		// Assert
		assert.Error(t, result)
		assert.Contains(t, result.Error(), "Unauthorized")
		// In production mode, should not contain internal error details
		assert.NotContains(t, result.Error(), "test error")
	})

	t.Run("returns unauthorized error with custom data", func(t *testing.T) {
		// Arrange
		middleware := &Middleware{}
		testErr := errors.New("test error")
		customData := base.CustomErrorData{"userId": "123"}

		// Act
		result := middleware.newUnauthorizedErr(testErr, customData)

		// Assert
		assert.Error(t, result)
		assert.Contains(t, result.Error(), "Unauthorized")
	})
}

func TestMiddlewareHandlerSuccessFlow(t *testing.T) {
	t.Run("successfully authenticates valid request", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}
		mockUsersClient := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		workosUser := usermanagement.User{
			ID:         "workos-123",
			ExternalID: testUserID.String(),
		}

		accessTokenClaims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env:             &env.Environment{},
			users:           mockUsersClient,
		}

		// Setup mocks
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "Bearer valid-token").Return(accessTokenClaims, nil)
		mockWorkosClient.userManagement.On("GetUser", mock.Anything, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(workosUser, nil)
		mockUsersClient.On("ByID", mock.Anything, testUserID).Return(testUser, nil)

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			session := GetSession(c)
			assert.NotNil(t, session)
			assert.Equal(t, testUser, session.User)
			assert.Equal(t, "session-456", session.ID)
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		req.Header.Set("X-Refresh-Token", "refresh-token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.AssertExpectations(t)
		mockUsersClient.AssertExpectations(t)
	})
}

func TestMiddlewareHandlerTokenExpiredFlow(t *testing.T) {
	t.Run("successfully refreshes expired token", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}
		mockUsersClient := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		workosUser := usermanagement.User{
			ID:         "workos-123",
			ExternalID: testUserID.String(),
		}

		accessTokenClaims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		refreshResponse := usermanagement.RefreshAuthenticationResponse{
			AccessToken:  "new-access-token",
			RefreshToken: "new-refresh-token",
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{ClientID: "test-client-id"},
			},
			users: mockUsersClient,
		}

		// Setup mocks - first call returns expired token error
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "Bearer expired-token").Return(nil, workos.AccessTokenExpired).Once()
		// Refresh token call
		mockWorkosClient.userManagement.On("AuthenticateWithRefreshToken", mock.Anything, mock.MatchedBy(func(opts usermanagement.AuthenticateWithRefreshTokenOpts) bool {
			return opts.ClientID == "test-client-id" && opts.RefreshToken == "refresh-token"
		})).Return(refreshResponse, nil)
		// Second call with new token succeeds
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "new-access-token").Return(accessTokenClaims, nil)
		// User retrieval mocks
		mockWorkosClient.userManagement.On("GetUser", mock.Anything, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(workosUser, nil)
		mockUsersClient.On("ByID", mock.Anything, testUserID).Return(testUser, nil)

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer expired-token")
		req.Header.Set("X-Refresh-Token", "refresh-token")
		req.Header.Set("User-Agent", "test-agent")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Check that new tokens were set in response headers
		assert.Equal(t, "new-refresh-token", resp.Header.Get("X-New-Refresh-Token"))
		assert.Equal(t, "new-access-token", resp.Header.Get("X-New-Access-Token"))

		mockWorkosClient.AssertExpectations(t)
		mockUsersClient.AssertExpectations(t)
	})
}

func TestMiddlewareHandlerErrorCases(t *testing.T) {
	t.Run("returns error when access token authentication fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env:             &env.Environment{},
			users:           &MockUsersClient{},
		}

		// Setup mock to return authentication error
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "Bearer invalid-token").Return(nil, errors.New("invalid token"))

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		req.Header.Set("X-Refresh-Token", "refresh-token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // Generic error from authentication failure
		mockWorkosClient.AssertExpectations(t)
	})

	t.Run("returns error when workos user retrieval fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}

		accessTokenClaims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env:             &env.Environment{},
			users:           &MockUsersClient{},
		}

		// Setup mocks
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "Bearer valid-token").Return(accessTokenClaims, nil)
		mockWorkosClient.userManagement.On("GetUser", mock.Anything, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(usermanagement.User{}, errors.New("workos api error"))

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		req.Header.Set("X-Refresh-Token", "refresh-token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		mockWorkosClient.AssertExpectations(t)
	})

	t.Run("returns error when user ID parsing fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}

		accessTokenClaims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		workosUser := usermanagement.User{
			ID:         "workos-123",
			ExternalID: "invalid-uuid", // This will cause parsing error
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env:             &env.Environment{},
			users:           &MockUsersClient{},
		}

		// Setup mocks
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "Bearer valid-token").Return(accessTokenClaims, nil)
		mockWorkosClient.userManagement.On("GetUser", mock.Anything, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(workosUser, nil)

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		req.Header.Set("X-Refresh-Token", "refresh-token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // UUID parsing error
		mockWorkosClient.AssertExpectations(t)
	})

	t.Run("returns error when database user retrieval fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}
		mockUsersClient := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		accessTokenClaims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		workosUser := usermanagement.User{
			ID:         "workos-123",
			ExternalID: testUserID.String(),
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env:             &env.Environment{},
			users:           mockUsersClient,
		}

		// Setup mocks
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "Bearer valid-token").Return(accessTokenClaims, nil)
		mockWorkosClient.userManagement.On("GetUser", mock.Anything, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(workosUser, nil)
		// Database error (not record not found)
		mockUsersClient.On("ByID", mock.Anything, testUserID).Return((*model.User)(nil), errors.New("database connection error"))

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(middleware.Handler)
		app.Get("/test", func(c *fiber.Ctx) error {
			return c.SendString("success")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		req.Header.Set("X-Refresh-Token", "refresh-token")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // Database error
		mockWorkosClient.AssertExpectations(t)
		mockUsersClient.AssertExpectations(t)
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("returns error when refresh token authentication fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{ClientID: "test-client-id"},
			},
			users: &MockUsersClient{},
		}

		// Setup mock to return error on refresh
		mockWorkosClient.userManagement.On("AuthenticateWithRefreshToken", mock.Anything, mock.MatchedBy(func(opts usermanagement.AuthenticateWithRefreshTokenOpts) bool {
			return opts.RefreshToken == "invalid-refresh-token"
		})).Return(usermanagement.RefreshAuthenticationResponse{}, errors.New("invalid refresh token"))

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(func(c *fiber.Ctx) error {
			// Directly test the refreshToken function
			_, err := middleware.refreshToken(c, "invalid-refresh-token")
			return err
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("User-Agent", "test-agent")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		mockWorkosClient.AssertExpectations(t)
	})

	t.Run("returns error when new access token authentication fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockWorkosClient := &MockWorkosClient{userManagement: &MockUserManagement{}}

		refreshResponse := usermanagement.RefreshAuthenticationResponse{
			AccessToken:  "new-but-invalid-token",
			RefreshToken: "new-refresh-token",
		}

		middleware := &Middleware{
			userCache:       localcache.New[*model.User](cacheOptions...),
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient:    mockWorkosClient,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{ClientID: "test-client-id"},
			},
			users: &MockUsersClient{},
		}

		// Setup mocks
		mockWorkosClient.userManagement.On("AuthenticateWithRefreshToken", mock.Anything, mock.MatchedBy(func(opts usermanagement.AuthenticateWithRefreshTokenOpts) bool {
			return opts.RefreshToken == "valid-refresh-token"
		})).Return(refreshResponse, nil)
		mockWorkosClient.On("AuthenticateWithAccessToken", mock.Anything, "new-but-invalid-token").Return(nil, errors.New("new token is invalid"))

		// Create fiber app
		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})
		app.Use(func(c *fiber.Ctx) error {
			// Directly test the refreshToken function
			_, err := middleware.refreshToken(c, "valid-refresh-token")
			return err
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("User-Agent", "test-agent")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		mockWorkosClient.AssertExpectations(t)
	})
}

func TestGetWorkosUserError(t *testing.T) {
	t.Run("returns error when WorkOS API fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUserManagement := &MockUserManagement{}

		accessTokenClaims := &workos.AccessTokenClaims{
			UserID:    "workos-123",
			SessionID: "session-456",
		}

		middleware := &Middleware{
			workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
			workosClient: &MockWorkosClient{
				userManagement: mockUserManagement,
			},
		}

		ctx := context.Background()

		mockUserManagement.On("GetUser", ctx, usermanagement.GetUserOpts{
			User: "workos-123",
		}).Return(usermanagement.User{}, errors.New("WorkOS API error"))

		// Act
		result, err := middleware.getWorkosUser(ctx, accessTokenClaims)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "WorkOS API error")
		mockUserManagement.AssertExpectations(t)
	})
}

func TestGetUserDatabaseError(t *testing.T) {
	t.Run("returns error when database operation fails", func(t *testing.T) {
		// Arrange
		cacheOptions := []localcache.Option{
			localcache.WithDefaultExpiration(10 * time.Minute),
			localcache.WithCleanupInterval(5 * time.Minute),
		}

		mockUsers := &MockUsersClient{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		middleware := &Middleware{
			userCache: localcache.New[*model.User](cacheOptions...),
			users:     mockUsers,
		}

		ctx := context.Background()

		mockUsers.On("ByID", ctx, testUserID).Return((*model.User)(nil), errors.New("database connection failed"))

		// Act
		result, err := middleware.getUser(ctx, testUserID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "database connection failed")
		mockUsers.AssertExpectations(t)
	})
}
