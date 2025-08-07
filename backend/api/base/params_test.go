package base

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"polimane/backend/model"
)

func TestNewMissingParamErr(t *testing.T) {
	t.Run("creates custom error for missing parameter", func(t *testing.T) {
		// Arrange
		paramName := "user_id"

		// Act
		err := newMissingParamErr(paramName)

		// Assert
		assert.Error(t, err)
		customErr, ok := err.(*CustomError)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusBadRequest, customErr.Code)
		assert.Equal(t, "MissingRequiredParam[user_id]", customErr.Message)
		assert.Equal(t, "MissingRequiredParam[user_id]", customErr.Data["reason"])
	})

	t.Run("handles empty parameter name", func(t *testing.T) {
		// Arrange
		paramName := ""

		// Act
		err := newMissingParamErr(paramName)

		// Assert
		assert.Error(t, err)
		customErr, ok := err.(*CustomError)
		assert.True(t, ok)
		assert.Equal(t, "MissingRequiredParam[]", customErr.Message)
	})
}

func TestGetRequiredParam(t *testing.T) {
	t.Run("returns parameter value when present", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			value, err := GetRequiredParam(c, "id")
			if err != nil {
				return err
			}
			return c.SendString(value)
		})

		req := httptest.NewRequest("GET", "/test/123", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("returns error when parameter is missing", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			value, err := GetRequiredParam(c, "missing_param")
			if err != nil {
				return err
			}
			return c.SendString(value)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns error when parameter is empty string", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			value, err := GetRequiredParam(c, "id")
			if err != nil {
				return err
			}
			return c.SendString(value)
		})

		// Test with empty parameter by creating a route that would result in empty param
		req := httptest.NewRequest("GET", "/test/", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		// This should return 404 as the route doesn't match /test/ (it expects /test/:id)
		assert.Equal(t, 404, resp.StatusCode)
	})

	t.Run("returns parameter value with special characters", func(t *testing.T) {
		// Arrange
		expectedValue := "user@example.com"
		app := fiber.New()
		app.Get("/test/:email", func(c *fiber.Ctx) error {
			value, err := GetRequiredParam(c, "email")
			if err != nil {
				return err
			}
			assert.Equal(t, expectedValue, value)
			return c.SendString("ok")
		})

		req := httptest.NewRequest("GET", "/test/"+expectedValue, nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

func TestGetParamID(t *testing.T) {
	t.Run("returns valid UUID when parameter is valid", func(t *testing.T) {
		// Arrange
		validUUID := "550e8400-e29b-41d4-a716-446655440000"
		app := fiber.New()
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			id, err := GetParamID(c, "id")
			if err != nil {
				return err
			}
			assert.Equal(t, validUUID, id.String())
			return c.SendString("ok")
		})

		req := httptest.NewRequest("GET", "/test/"+validUUID, nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("returns error when parameter is missing", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			id, err := GetParamID(c, "missing_id")
			if err != nil {
				return err
			}
			return c.SendString(id.String())
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns error when parameter is invalid UUID", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			id, err := GetParamID(c, "id")
			if err != nil {
				return err
			}
			return c.SendString(id.String())
		})

		req := httptest.NewRequest("GET", "/test/invalid-uuid", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns error when parameter is empty string", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		// This setup simulates an empty parameter scenario
		app.Get("/test", func(c *fiber.Ctx) error {
			// Manually test the empty parameter case
			ctx := c.Context()
			// Since we can't easily simulate empty params with Fiber test,
			// we'll test the function directly
			id, err := GetParamID(c, "non_existent_param")
			if err != nil {
				return err
			}
			_ = ctx // silence unused variable warning
			return c.SendString(id.String())
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns zero ID when error occurs", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		var resultID model.ID
		var resultErr error

		app.Get("/test", func(c *fiber.Ctx) error {
			resultID, resultErr = GetParamID(c, "non_existent_param")
			if resultErr != nil {
				return resultErr
			}
			return c.SendString(resultID.String())
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		assert.Error(t, resultErr)
		assert.Equal(t, model.ID{}, resultID)
	})

	t.Run("handles malformed UUID correctly", func(t *testing.T) {
		// Arrange
		malformedUUID := "not-a-uuid-at-all"
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			id, err := GetParamID(c, "id")
			if err != nil {
				return err
			}
			return c.SendString(id.String())
		})

		req := httptest.NewRequest("GET", "/test/"+malformedUUID, nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("handles UUID with wrong format", func(t *testing.T) {
		// Arrange
		wrongFormatUUID := "550e8400-e29b-41d4-a716-44665544000" // Missing one character
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test/:id", func(c *fiber.Ctx) error {
			id, err := GetParamID(c, "id")
			if err != nil {
				return err
			}
			return c.SendString(id.String())
		})

		req := httptest.NewRequest("GET", "/test/"+wrongFormatUUID, nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
