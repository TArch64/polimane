package users

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func TestController_apiGet(t *testing.T) {
	t.Run("returns current user data successfully", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		mockSignals := NewMockSignalsContainer()
		testUser := createTestUser()
		testWorkosUser := createTestWorkosUser()

		controller := &Controller{
			workosClient: mockWorkosClient,
			signals:      mockSignals,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/users/current", func(c *fiber.Ctx) error {
			// Set up session with both User and WorkosUser
			auth.SetSession(c, &auth.UserSession{
				ID:         "session-123",
				User:       testUser,
				WorkosUser: testWorkosUser,
			})
			return controller.apiGet(c)
		})

		req := httptest.NewRequest("GET", "/users/current", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseUser currentUser
		err = json.Unmarshal(body, &responseUser)
		assert.NoError(t, err)

		assert.Equal(t, testUser.ID, responseUser.ID)
		assert.Equal(t, testWorkosUser.FirstName, responseUser.FirstName)
		assert.Equal(t, testWorkosUser.LastName, responseUser.LastName)
		assert.Equal(t, testWorkosUser.Email, responseUser.Email)
		assert.Equal(t, testWorkosUser.EmailVerified, responseUser.EmailVerified)
		assert.Equal(t, testWorkosUser.ProfilePictureURL, responseUser.ProfilePictureURL)
	})

}
