package users

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func TestController_apiAuthFactorsInit(t *testing.T) {
	t.Run("successfully initializes auth factor", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		// Mock response with empty values - the actual structure is complex
		mockEnrollResponse := usermanagement.EnrollAuthFactorResponse{}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("EnrollAuthFactor",
			mock.Anything,
			usermanagement.EnrollAuthFactorOpts{
				User:       testWorkosUser.ID,
				Type:       mfa.TOTP,
				TOTPIssuer: "Polimane",
				TOTPUser:   testWorkosUser.Email,
			}).Return(mockEnrollResponse, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors/init", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiAuthFactorsInit(c)
		})

		req := httptest.NewRequest("POST", "/users/current/auth-factors/init", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify JSON response structure without checking exact values
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseBody map[string]interface{}
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		// Check that the expected JSON fields exist
		assert.Contains(t, responseBody, "challengeId")
		assert.Contains(t, responseBody, "qrCode")
		assert.Contains(t, responseBody, "secret")
		assert.Contains(t, responseBody, "uri")

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("uses correct TOTP configuration parameters", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()
		testWorkosUser.ID = "specific-workos-id"
		testWorkosUser.Email = "specific@test.com"

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockEnrollResponse := usermanagement.EnrollAuthFactorResponse{}

		// Verify that the correct parameters are used for TOTP enrollment
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("EnrollAuthFactor",
			mock.Anything,
			usermanagement.EnrollAuthFactorOpts{
				User:       "specific-workos-id", // Should use WorkOS user ID
				Type:       mfa.TOTP,             // Should be TOTP type
				TOTPIssuer: "Polimane",           // Should always be "Polimane"
				TOTPUser:   "specific@test.com",  // Should use WorkOS user email
			}).Return(mockEnrollResponse, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors/init", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiAuthFactorsInit(c)
		})

		req := httptest.NewRequest("POST", "/users/current/auth-factors/init", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("handles WorkOS enrollment error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		enrollError := &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "User already has maximum auth factors",
		}

		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("EnrollAuthFactor",
			mock.Anything,
			mock.Anything).Return(usermanagement.EnrollAuthFactorResponse{}, enrollError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors/init", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiAuthFactorsInit(c)
		})

		req := httptest.NewRequest("POST", "/users/current/auth-factors/init", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("accesses session WorkOS user correctly", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockEnrollResponse := usermanagement.EnrollAuthFactorResponse{}

		// Verify that session.WorkosUser is accessed for both ID and Email
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).On("EnrollAuthFactor",
			mock.Anything,
			mock.MatchedBy(func(opts usermanagement.EnrollAuthFactorOpts) bool {
				return opts.User == testWorkosUser.ID &&
					opts.TOTPUser == testWorkosUser.Email &&
					opts.TOTPIssuer == "Polimane" &&
					opts.Type == mfa.TOTP
			})).Return(mockEnrollResponse, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors/init", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiAuthFactorsInit(c)
		})

		req := httptest.NewRequest("POST", "/users/current/auth-factors/init", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement().(*MockWorkosUserManagement).AssertExpectations(t)
	})
}

func TestNewAuthFactorResponse(t *testing.T) {
	t.Run("has correct JSON structure", func(t *testing.T) {
		// Arrange
		response := newAuthFactorResponse{
			ChallengeID: "challenge-123",
			QRCode:      "qr-code-data",
			Secret:      "secret-key",
			URI:         "otpauth://totp/example",
		}

		// Act
		jsonBytes, err := json.Marshal(response)

		// Assert
		assert.NoError(t, err)

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonBytes, &unmarshaled)
		assert.NoError(t, err)

		assert.Equal(t, "challenge-123", unmarshaled["challengeId"])
		assert.Equal(t, "qr-code-data", unmarshaled["qrCode"])
		assert.Equal(t, "secret-key", unmarshaled["secret"])
		assert.Equal(t, "otpauth://totp/example", unmarshaled["uri"])
	})
}
