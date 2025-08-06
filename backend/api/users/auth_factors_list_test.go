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

func TestController_apiListAuthFactors(t *testing.T) {
	t.Run("successfully returns list of auth factors", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockFactors := usermanagement.ListAuthFactorsResponse{
			Data: []mfa.Factor{
				{
					ID:        "factor-123",
					CreatedAt: "2023-01-01T12:00:00Z",
				},
				{
					ID:        "factor-456",
					CreatedAt: "2023-01-02T13:00:00Z",
				},
			},
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("ListAuthFactors",
			mock.Anything,
			usermanagement.ListAuthFactorsOpts{
				User: testUser.WorkosID,
			}).Return(mockFactors, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiListAuthFactors(c)
		})

		req := httptest.NewRequest("GET", "/users/current/auth-factors", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseFactors []authFactorListItem
		err = json.Unmarshal(body, &responseFactors)
		assert.NoError(t, err)

		assert.Len(t, responseFactors, 2)
		assert.Equal(t, "factor-123", responseFactors[0].Id)
		assert.Equal(t, "2023-01-01T12:00:00Z", responseFactors[0].CreatedAt)
		assert.Equal(t, "factor-456", responseFactors[1].Id)
		assert.Equal(t, "2023-01-02T13:00:00Z", responseFactors[1].CreatedAt)

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("returns empty array when no auth factors exist", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockFactors := usermanagement.ListAuthFactorsResponse{
			Data: []mfa.Factor{}, // Empty array
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("ListAuthFactors",
			mock.Anything,
			usermanagement.ListAuthFactorsOpts{
				User: testUser.WorkosID,
			}).Return(mockFactors, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiListAuthFactors(c)
		})

		req := httptest.NewRequest("GET", "/users/current/auth-factors", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseFactors []authFactorListItem
		err = json.Unmarshal(body, &responseFactors)
		assert.NoError(t, err)

		assert.Len(t, responseFactors, 0)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("handles WorkOS error", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		workosError := &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Service unavailable",
		}

		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("ListAuthFactors",
			mock.Anything,
			usermanagement.ListAuthFactorsOpts{
				User: testUser.WorkosID,
			}).Return(usermanagement.ListAuthFactorsResponse{}, workosError)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiListAuthFactors(c)
		})

		req := httptest.NewRequest("GET", "/users/current/auth-factors", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})

	t.Run("uses correct user ID from session", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		testUser := createTestUser()
		testUser.WorkosID = "specific-workos-user-id"

		controller := &Controller{
			workosClient: mockWorkosClient,
		}

		mockFactors := usermanagement.ListAuthFactorsResponse{
			Data: []mfa.Factor{},
		}

		// Verify that the specific WorkOS user ID is used
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).On("ListAuthFactors",
			mock.Anything,
			usermanagement.ListAuthFactorsOpts{
				User: "specific-workos-user-id", // Must match the user's WorkOS ID
			}).Return(mockFactors, nil)

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/users/current/auth-factors", func(c *fiber.Ctx) error {
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiListAuthFactors(c)
		})

		req := httptest.NewRequest("GET", "/users/current/auth-factors", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockWorkosClient.UserManagement.(*MockWorkosUserManagement).AssertExpectations(t)
	})
}

func TestAuthFactorListItem(t *testing.T) {
	t.Run("has correct JSON structure", func(t *testing.T) {
		// Arrange
		item := authFactorListItem{
			Id:        "factor-123",
			CreatedAt: "2023-01-01T12:00:00Z",
		}

		// Act
		jsonBytes, err := json.Marshal(item)

		// Assert
		assert.NoError(t, err)

		var unmarshaled map[string]interface{}
		err = json.Unmarshal(jsonBytes, &unmarshaled)
		assert.NoError(t, err)

		assert.Equal(t, "factor-123", unmarshaled["id"])
		assert.Equal(t, "2023-01-01T12:00:00Z", unmarshaled["createdAt"])
	})
}
