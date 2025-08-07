package base

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type testRequestBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func TestParseBody(t *testing.T) {
	// Initialize validator for tests
	InitValidator()

	t.Run("successfully parses valid JSON body", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Post("/test", func(c *fiber.Ctx) error {
			var body testRequestBody
			err := ParseBody(c, &body)
			if err != nil {
				return err
			}
			return c.JSON(body)
		})

		jsonBody := `{"name":"John Doe","email":"john@example.com"}`
		req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("returns error for invalid JSON", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Post("/test", func(c *fiber.Ctx) error {
			var body testRequestBody
			err := ParseBody(c, &body)
			if err != nil {
				return err
			}
			return c.JSON(body)
		})

		invalidJson := `{"name":"John Doe","email":}`
		req := httptest.NewRequest("POST", "/test", strings.NewReader(invalidJson))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // JSON parsing error
	})

	t.Run("returns validation error for missing required fields", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Post("/test", func(c *fiber.Ctx) error {
			var body testRequestBody
			err := ParseBody(c, &body)
			if err != nil {
				return err
			}
			return c.JSON(body)
		})

		jsonBody := `{"email":"john@example.com"}`
		req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode) // Validation error
	})

	t.Run("returns validation error for invalid email format", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Post("/test", func(c *fiber.Ctx) error {
			var body testRequestBody
			err := ParseBody(c, &body)
			if err != nil {
				return err
			}
			return c.JSON(body)
		})

		jsonBody := `{"name":"John Doe","email":"invalid-email"}`
		req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode) // Validation error
	})

	t.Run("handles empty request body", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Post("/test", func(c *fiber.Ctx) error {
			var body testRequestBody
			err := ParseBody(c, &body)
			if err != nil {
				return err
			}
			return c.JSON(body)
		})

		req := httptest.NewRequest("POST", "/test", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // JSON parsing error for empty body
	})
}
