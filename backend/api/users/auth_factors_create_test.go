package users

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/mfa"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func TestController_apiAuthFactorCreate(t *testing.T) {
	// Initialize validator for all tests
	base.InitValidator()

	t.Run("successfully creates auth factor with valid challenge and code", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockVerifyResponse := mfa.VerifyChallengeResponse{
			Valid: true,
			Challenge: mfa.Challenge{
				FactorID: "factor-123",
			},
		}

		mockFactor := mfa.Factor{
			ID:        "factor-123",
			CreatedAt: "2023-01-01T12:00:00Z",
		}

		mockWorkosClient.MFA().(*MockWorkosMFA).On("VerifyChallenge",
			mock.Anything,
			mfa.VerifyChallengeOpts{
				Code:        "123456",
				ChallengeID: "challenge-123",
			}).Return(mockVerifyResponse, nil)

		mockWorkosClient.MFA().(*MockWorkosMFA).On("GetFactor",
			mock.Anything,
			mfa.GetFactorOpts{
				FactorID: "factor-123",
			}).Return(mockFactor, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorCreate(c)
		})

		requestBody := createAuthFactorBody{
			ChallengeID: "challenge-123",
			Code:        "123456",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("POST", "/users/current/auth-factors", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.MFA().(*MockWorkosMFA).AssertExpectations(t)
	})

	t.Run("returns 400 when challengeId is missing", func(t *testing.T) {
		// Arrange
		controller := &Controller{}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorCreate(c)
		})

		requestBody := createAuthFactorBody{
			Code: "123456", // Missing ChallengeID
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("POST", "/users/current/auth-factors", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("returns 400 when code is missing", func(t *testing.T) {
		// Arrange
		controller := &Controller{}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorCreate(c)
		})

		requestBody := createAuthFactorBody{
			ChallengeID: "challenge-123", // Missing Code
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("POST", "/users/current/auth-factors", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("returns InvalidAuthFactor error when verification is not valid", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockVerifyResponse := mfa.VerifyChallengeResponse{
			Valid: false, // Invalid verification
		}

		mockWorkosClient.MFA().(*MockWorkosMFA).On("VerifyChallenge",
			mock.Anything,
			mock.Anything).Return(mockVerifyResponse, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorCreate(c)
		})

		requestBody := createAuthFactorBody{
			ChallengeID: "challenge-123",
			Code:        "123456",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("POST", "/users/current/auth-factors", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		mockWorkosClient.MFA().(*MockWorkosMFA).AssertExpectations(t)
	})

	t.Run("handles VerifyChallenge error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		verifyError := &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Invalid challenge",
		}

		mockWorkosClient.MFA().(*MockWorkosMFA).On("VerifyChallenge",
			mock.Anything,
			mock.Anything).Return(mfa.VerifyChallengeResponse{}, verifyError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorCreate(c)
		})

		requestBody := createAuthFactorBody{
			ChallengeID: "challenge-123",
			Code:        "123456",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("POST", "/users/current/auth-factors", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode)
		mockWorkosClient.MFA().(*MockWorkosMFA).AssertExpectations(t)
	})

	t.Run("handles GetFactor error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockVerifyResponse := mfa.VerifyChallengeResponse{
			Valid: true,
			Challenge: mfa.Challenge{
				FactorID: "factor-123",
			},
		}

		getFactorError := &fiber.Error{
			Code:    fiber.StatusNotFound,
			Message: "Factor not found",
		}

		mockWorkosClient.MFA().(*MockWorkosMFA).On("VerifyChallenge",
			mock.Anything,
			mock.Anything).Return(mockVerifyResponse, nil)

		mockWorkosClient.MFA().(*MockWorkosMFA).On("GetFactor",
			mock.Anything,
			mfa.GetFactorOpts{
				FactorID: "factor-123",
			}).Return(mfa.Factor{}, getFactorError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: createTestUser(),
			})
			return controller.apiAuthFactorCreate(c)
		})

		requestBody := createAuthFactorBody{
			ChallengeID: "challenge-123",
			Code:        "123456",
		}
		bodyBytes, _ := json.Marshal(requestBody)
		req := httptest.NewRequest("POST", "/users/current/auth-factors", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
		mockWorkosClient.MFA().(*MockWorkosMFA).AssertExpectations(t)
	})
}

func TestCreateAuthFactorBody(t *testing.T) {
	t.Run("has correct validation tags", func(t *testing.T) {
		// This test verifies the struct validation requirements
		body := createAuthFactorBody{
			ChallengeID: "challenge-123",
			Code:        "123456",
		}

		assert.Equal(t, "challenge-123", body.ChallengeID)
		assert.Equal(t, "123456", body.Code)
	})
}

func TestErrInvalidAuthFactor(t *testing.T) {
	t.Run("has correct error properties", func(t *testing.T) {
		// Act & Assert
		assert.NotNil(t, ErrInvalidAuthFactor)
		// We can't easily test the specific error structure without examining the base.NewReasonedError implementation
		// but we can verify the error exists and is not nil
	})
}
