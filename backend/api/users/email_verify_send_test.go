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

func TestController_apiEmailVerifyRetry(t *testing.T) {
	// Initialize validator for all tests
	base.InitValidator()

	t.Run("successfully sends verification email retry", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything,
			usermanagement.SendVerificationEmailOpts{
				User: testUser.WorkosID,
			}).Return(usermanagement.UserResponse{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/email/verify/retry", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiEmailVerifyRetry(c)
		})

		req := httptest.NewRequest("POST", "/users/current/email/verify/retry", nil)

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

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		workosError := &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Email already verified",
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything,
			usermanagement.SendVerificationEmailOpts{
				User: testUser.WorkosID,
			}).Return(usermanagement.UserResponse{}, workosError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/email/verify/retry", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiEmailVerifyRetry(c)
		})

		req := httptest.NewRequest("POST", "/users/current/email/verify/retry", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("uses user ID from session", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()
		testUser.WorkosID = "specific-workos-id"

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		// Verify that the specific WorkOS user ID is used
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything,
			usermanagement.SendVerificationEmailOpts{
				User: "specific-workos-id", // Must match the user's WorkOS ID
			}).Return(usermanagement.UserResponse{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/email/verify/retry", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiEmailVerifyRetry(c)
		})

		req := httptest.NewRequest("POST", "/users/current/email/verify/retry", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})
}

func TestSendEmailVerification(t *testing.T) {
	t.Run("calls WorkOS SendVerificationEmail with correct parameters", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		expectedOpts := usermanagement.SendVerificationEmailOpts{
			User: "user-123",
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything, expectedOpts).Return(usermanagement.UserResponse{}, nil)

		// Act
		err := controller.sendEmailVerification(nil, "user-123")

		// Assert
		assert.NoError(t, err)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("returns error when WorkOS call fails", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		workosError := &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Service unavailable",
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything, mock.Anything).Return(usermanagement.UserResponse{}, workosError)

		// Act
		err := controller.sendEmailVerification(nil, "user-123")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, workosError, err)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})
}
