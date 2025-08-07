package base

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewSuccessResponse(t *testing.T) {
	t.Run("returns success JSON response", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test", func(c *fiber.Ctx) error {
			return NewSuccessResponse(c)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		// Check response body
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response map[string]interface{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, true, response["success"])
		assert.Len(t, response, 1) // Should only contain "success" field
	})

	t.Run("sets correct content type", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test", func(c *fiber.Ctx) error {
			return NewSuccessResponse(c)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	})

	t.Run("can be used in middleware chain", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Add a middleware that sets a header
		app.Use(func(c *fiber.Ctx) error {
			c.Set("X-Test-Header", "test-value")
			return c.Next()
		})

		app.Get("/test", func(c *fiber.Ctx) error {
			return NewSuccessResponse(c)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "test-value", resp.Header.Get("X-Test-Header"))

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response map[string]interface{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, true, response["success"])
	})

	t.Run("maintains status code when called", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Post("/test", func(c *fiber.Ctx) error {
			// Default POST status would be 200, but we can verify it stays 200
			return NewSuccessResponse(c)
		})

		req := httptest.NewRequest("POST", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("returns proper JSON structure", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test", func(c *fiber.Ctx) error {
			return NewSuccessResponse(c)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		// Verify it's valid JSON
		var response fiber.Map
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		// Verify the structure matches fiber.Map{"success": true}
		expected := fiber.Map{"success": true}
		assert.Equal(t, expected, response)
	})

	t.Run("handles context properly", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			// Verify that the context is still functional after success response
			id := c.Params("id")
			assert.Equal(t, "123", id)
			return NewSuccessResponse(c)
		})

		req := httptest.NewRequest("GET", "/test/123", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
