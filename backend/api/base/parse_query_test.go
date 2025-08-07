package base

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type testQueryStruct struct {
	Name string `query:"name"`
	Age  int    `query:"age"`
}

type testValidatedQuery struct {
	Email string `query:"email" validate:"required,email"`
}

func TestParseQuery(t *testing.T) {
	// Initialize validator for tests
	InitValidator()

	t.Run("successfully parses valid query parameters", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test", func(c *fiber.Ctx) error {
			var query testQueryStruct
			err := ParseQuery(c, &query)
			if err != nil {
				return err
			}
			return c.JSON(query)
		})

		req := httptest.NewRequest("GET", "/test?name=John&age=30", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("returns error when query parsing fails", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			var query testQueryStruct
			err := ParseQuery(c, &query)
			if err != nil {
				return err
			}
			return c.JSON(query)
		})

		// Invalid integer value should cause query parser to fail
		req := httptest.NewRequest("GET", "/test?name=John&age=not-a-number", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // Query parser error
	})

	t.Run("returns error when validation fails", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			var query testValidatedQuery
			err := ParseQuery(c, &query)
			if err != nil {
				return err
			}
			return c.JSON(query)
		})

		// Missing required field should cause validation to fail
		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode) // Validation error
	})

	t.Run("returns nil when both parsing and validation succeed", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test", func(c *fiber.Ctx) error {
			var query testValidatedQuery
			err := ParseQuery(c, &query)
			if err != nil {
				return err
			}
			return c.JSON(query)
		})

		req := httptest.NewRequest("GET", "/test?email=test@example.com", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("handles empty query string", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test", func(c *fiber.Ctx) error {
			var query testQueryStruct
			err := ParseQuery(c, &query)
			if err != nil {
				return err
			}
			return c.JSON(query)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode) // Should succeed with empty values
	})
}
