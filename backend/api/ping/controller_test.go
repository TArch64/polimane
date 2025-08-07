package ping

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	t.Run("returns Controller instance", func(t *testing.T) {
		// Act
		controller := Provider()

		// Assert
		assert.NotNil(t, controller)
		assert.IsType(t, &Controller{}, controller)
	})
}

func TestController_Public(t *testing.T) {
	t.Run("does not add any public routes", func(t *testing.T) {
		// Arrange
		controller := &Controller{}
		app := fiber.New()

		// Act
		controller.Public(app)

		// Assert - no routes added, should return 404
		req := httptest.NewRequest("GET", "/ping", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
	})
}

func TestController_Private(t *testing.T) {
	t.Run("adds ping route to private group", func(t *testing.T) {
		// Arrange
		controller := &Controller{}
		app := fiber.New()
		privateGroup := app.Group("private")

		// Act
		controller.Private(privateGroup)

		// Test the ping endpoint
		req := httptest.NewRequest("GET", "/private/ping", nil)
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
