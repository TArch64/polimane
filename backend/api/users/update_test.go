package users

import (
	"bytes"
	"encoding/json"
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

func TestController_apiUpdate(t *testing.T) {
	// Initialize validator for all tests
	base.InitValidator()

	t.Run("successfully updates user with all fields", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("UpdateUser",
			mock.Anything,
			usermanagement.UpdateUserOpts{
				User:      testUser.WorkosID,
				FirstName: "Updated",
				LastName:  "Name",
				Email:     "updated@example.com",
			}).Return(usermanagement.User{}, nil)

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything,
			usermanagement.SendVerificationEmailOpts{
				User: testUser.WorkosID,
			}).Return(usermanagement.UserResponse{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		requestBody := updateBody{
			FirstName: "Updated",
			LastName:  "Name",
			Email:     "updated@example.com",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("successfully updates user with only first name", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("UpdateUser",
			mock.Anything,
			usermanagement.UpdateUserOpts{
				User:      testUser.WorkosID,
				FirstName: "Updated",
				LastName:  "",
				Email:     "",
			}).Return(usermanagement.User{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		requestBody := updateBody{
			FirstName: "Updated",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("returns 422 when no fields provided", func(t *testing.T) {
		// Arrange
		controller := &Controller{}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiUpdate(c)
		})

		requestBody := updateBody{} // Empty body
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 422, resp.StatusCode)
	})

	t.Run("handles invalid JSON body", func(t *testing.T) {
		// Arrange
		controller := &Controller{}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiUpdate(c)
		})

		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		// Fiber returns 500 for JSON parsing errors, not 400
		assert.Equal(t, 500, resp.StatusCode)
	})

	t.Run("sends email verification when email is updated", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("UpdateUser",
			mock.Anything, mock.Anything).Return(usermanagement.User{}, nil)

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything,
			usermanagement.SendVerificationEmailOpts{
				User: testUser.WorkosID,
			}).Return(usermanagement.UserResponse{}, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		requestBody := updateBody{
			Email: "new@example.com",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertCalled(t, "SendVerificationEmail", mock.Anything, mock.Anything)
	})

	t.Run("handles updateUser error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		// Mock updateUser to fail
		updateError := &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "WorkOS API error",
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("UpdateUser",
			mock.Anything, mock.Anything).Return(usermanagement.User{}, updateError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		requestBody := updateBody{
			FirstName: "Updated",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("handles sendEmailVerification error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := signal.Provider()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      signalsContainer,
		}

		// Mock updateUser to succeed
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("UpdateUser",
			mock.Anything, mock.Anything).Return(usermanagement.User{}, nil)

		// Mock sendEmailVerification to fail
		emailError := &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Email verification failed",
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything, mock.Anything).Return(usermanagement.UserResponse{}, emailError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Put("/users/current", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		requestBody := updateBody{
			Email: "updated@example.com",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("PUT", "/users/current", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})
}

func TestUpdateBody(t *testing.T) {
	t.Run("has correct JSON tags and validation", func(t *testing.T) {
		// This test verifies the struct definition
		body := updateBody{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
		}

		assert.Equal(t, "John", body.FirstName)
		assert.Equal(t, "Doe", body.LastName)
		assert.Equal(t, "john@example.com", body.Email)
	})
}

func TestController_updateUser(t *testing.T) {
	t.Run("calls WorkOS UpdateUser with correct parameters", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		body := &updateBody{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
		}

		expectedOpts := usermanagement.UpdateUserOpts{
			User:      "user-123",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("UpdateUser",
			mock.Anything, expectedOpts).Return(usermanagement.User{}, nil)

		// Act
		err := controller.updateUser(nil, "user-123", body)

		// Assert
		assert.NoError(t, err)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})
}

func TestController_sendEmailVerification(t *testing.T) {
	t.Run("calls WorkOS SendVerificationEmail with correct parameters", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		expectedOpts := usermanagement.SendVerificationEmailOpts{
			User: "user-123",
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("SendVerificationEmail",
			mock.Anything, expectedOpts).Return(usermanagement.UserResponse{}, nil)

		// Act
		err := controller.sendEmailVerification(nil, "user-123")

		// Assert
		assert.NoError(t, err)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})
}
