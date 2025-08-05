package auth

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

func TestApiLogout(t *testing.T) {
	t.Run("logs out successfully", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockSignal := &MockSignal[string]{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		session := &UserSession{
			ID:   "session-123",
			User: testUser,
			WorkosUser: &usermanagement.User{
				ID: "workos-user-123",
			},
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			signals: &signal.Container{
				InvalidateAuthCache: mockSignal,
			},
		}

		mockUserManagement.On("RevokeSession", mock.Anything, mock.MatchedBy(func(opts usermanagement.RevokeSessionOpts) bool {
			return opts.SessionID == "session-123"
		})).Return(nil)

		mockSignal.On("Emit", mock.Anything, "session-123").Return()

		// Create fiber app and request
		app := fiber.New()
		app.Post("/logout", func(c *fiber.Ctx) error {
			setSession(c, session)
			return controller.apiLogout(c)
		})
		req := httptest.NewRequest("POST", "/logout", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Check response body is success response
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, true, response["success"])

		mockUserManagement.AssertExpectations(t)
		mockSignal.AssertExpectations(t)
	})

	t.Run("handles WorkOS revoke session error", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		expectedError := assert.AnError

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		session := &UserSession{
			ID:   "session-456",
			User: testUser,
			WorkosUser: &usermanagement.User{
				ID: "workos-user-456",
			},
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			signals: &signal.Container{},
		}

		mockUserManagement.On("RevokeSession", mock.Anything, mock.Anything).Return(expectedError)

		// Create fiber app and request
		app := fiber.New()
		app.Post("/logout", func(c *fiber.Ctx) error {
			setSession(c, session)
			return controller.apiLogout(c)
		})
		req := httptest.NewRequest("POST", "/logout", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
	})

	t.Run("uses session ID from context", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockSignal := &MockSignal[string]{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		customSessionID := "custom-session-789"
		session := &UserSession{
			ID:   customSessionID,
			User: testUser,
			WorkosUser: &usermanagement.User{
				ID: "workos-user-789",
			},
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			signals: &signal.Container{
				InvalidateAuthCache: mockSignal,
			},
		}

		mockUserManagement.On("RevokeSession", mock.Anything, mock.MatchedBy(func(opts usermanagement.RevokeSessionOpts) bool {
			return opts.SessionID == customSessionID
		})).Return(nil)

		mockSignal.On("Emit", mock.Anything, customSessionID).Return()

		// Create fiber app and request
		app := fiber.New()
		app.Post("/logout", func(c *fiber.Ctx) error {
			setSession(c, session)
			return controller.apiLogout(c)
		})
		req := httptest.NewRequest("POST", "/logout", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockUserManagement.AssertExpectations(t)
		mockSignal.AssertExpectations(t)
	})

	t.Run("emits invalidate auth cache signal after successful revoke", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockSignal := &MockSignal[string]{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		session := &UserSession{
			ID:   "emit-test-session",
			User: testUser,
			WorkosUser: &usermanagement.User{
				ID: "workos-user-emit-test",
			},
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			signals: &signal.Container{
				InvalidateAuthCache: mockSignal,
			},
		}

		// Set up expectations - signal should be called after successful revoke
		mockUserManagement.On("RevokeSession", mock.Anything, mock.Anything).Return(nil)
		mockSignal.On("Emit", mock.Anything, "emit-test-session").Return()

		// Create fiber app and request
		app := fiber.New()
		app.Post("/logout", func(c *fiber.Ctx) error {
			setSession(c, session)
			return controller.apiLogout(c)
		})
		req := httptest.NewRequest("POST", "/logout", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify both WorkOS revoke and signal emit were called
		mockUserManagement.AssertExpectations(t)
		mockSignal.AssertExpectations(t)
	})

	t.Run("returns content type application/json", func(t *testing.T) {
		// Arrange
		mockUserManagement := &MockUserManagement{}
		mockSignal := &MockSignal[string]{}

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		testUser := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		session := &UserSession{
			ID:   "content-type-test",
			User: testUser,
			WorkosUser: &usermanagement.User{
				ID: "workos-user-content-type",
			},
		}

		controller := &Controller{
			workosClient: &workos.Client{
				UserManagement: mockUserManagement,
			},
			signals: &signal.Container{
				InvalidateAuthCache: mockSignal,
			},
		}

		mockUserManagement.On("RevokeSession", mock.Anything, mock.Anything).Return(nil)
		mockSignal.On("Emit", mock.Anything, "content-type-test").Return()

		// Create fiber app and request
		app := fiber.New()
		app.Post("/logout", func(c *fiber.Ctx) error {
			setSession(c, session)
			return controller.apiLogout(c)
		})
		req := httptest.NewRequest("POST", "/logout", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		mockUserManagement.AssertExpectations(t)
		mockSignal.AssertExpectations(t)
	})
}
