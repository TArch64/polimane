package users

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/signal"
)

func TestController_apiPasswordReset(t *testing.T) {
	// Initialize validator for all tests
	base.InitValidator()

	t.Run("successfully creates password reset", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("CreatePasswordReset",
			mock.Anything,
			usermanagement.CreatePasswordResetOpts{
				Email: testWorkosUser.Email,
			}).Return(usermanagement.PasswordReset{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/password/reset", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiPasswordReset(c)
		})

		req := httptest.NewRequest("POST", "/users/current/password/reset", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("handles WorkOS error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		workosError := &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid email",
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("CreatePasswordReset",
			mock.Anything,
			usermanagement.CreatePasswordResetOpts{
				Email: testWorkosUser.Email,
			}).Return(usermanagement.PasswordReset{}, workosError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/password/reset", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiPasswordReset(c)
		})

		req := httptest.NewRequest("POST", "/users/current/password/reset", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("uses email from WorkOS user in session", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()
		testWorkosUser.Email = "specific@test.com" // Set specific email

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		// Verify that the specific email from WorkOS user is used
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("CreatePasswordReset",
			mock.Anything,
			usermanagement.CreatePasswordResetOpts{
				Email: "specific@test.com", // Must match the WorkOS user email
			}).Return(usermanagement.PasswordReset{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/password/reset", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiPasswordReset(c)
		})

		req := httptest.NewRequest("POST", "/users/current/password/reset", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})
}
