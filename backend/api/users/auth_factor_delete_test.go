package users

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/mfa"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func TestController_apiAuthFactorDelete(t *testing.T) {
	t.Run("successfully deletes auth factor with valid factor ID", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockWorkosClient.MFA.(*MockWorkosMFA).On("DeleteFactor",
			mock.Anything,
			mfa.DeleteFactorOpts{
				FactorID: "factor-123",
			}).Return(nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Delete("/users/current/auth-factors/:factorId", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/users/current/auth-factors/factor-123", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.MFA.(*MockWorkosMFA).AssertExpectations(t)
	})

	t.Run("returns 400 when factor ID parameter is missing", func(t *testing.T) {
		// Arrange
		controller := &Controller{}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		// Route without the factorId parameter
		app.Delete("/users/current/auth-factors/", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/users/current/auth-factors/", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("returns 404 when factor ID parameter is empty", func(t *testing.T) {
		// Arrange
		controller := &Controller{}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Delete("/users/current/auth-factors/:factorId", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/users/current/auth-factors/", nil) // Empty factorId

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		// Fiber returns 404 when route doesn't match, not 400
		assert.Equal(t, 404, resp.StatusCode)
	})

	t.Run("handles WorkOS delete error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		deleteError := &fiber.Error{
			Code:    fiber.StatusNotFound,
			Message: "Factor not found",
		}

		mockWorkosClient.MFA.(*MockWorkosMFA).On("DeleteFactor",
			mock.Anything,
			mfa.DeleteFactorOpts{
				FactorID: "factor-123",
			}).Return(deleteError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Delete("/users/current/auth-factors/:factorId", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/users/current/auth-factors/factor-123", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
		mockWorkosClient.MFA.(*MockWorkosMFA).AssertExpectations(t)
	})

	t.Run("uses correct factor ID from URL parameter", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		// Verify that the specific factor ID from the URL is used
		mockWorkosClient.MFA.(*MockWorkosMFA).On("DeleteFactor",
			mock.Anything,
			mfa.DeleteFactorOpts{
				FactorID: "specific-factor-id-456", // Must match the URL parameter
			}).Return(nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Delete("/users/current/auth-factors/:factorId", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/users/current/auth-factors/specific-factor-id-456", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.MFA.(*MockWorkosMFA).AssertExpectations(t)
	})

	t.Run("handles unauthorized access error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		unauthorizedError := &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Factor belongs to different user",
		}

		mockWorkosClient.MFA.(*MockWorkosMFA).On("DeleteFactor",
			mock.Anything,
			mock.Anything).Return(unauthorizedError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Delete("/users/current/auth-factors/:factorId", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/users/current/auth-factors/factor-123", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode)
		mockWorkosClient.MFA.(*MockWorkosMFA).AssertExpectations(t)
	})
}
