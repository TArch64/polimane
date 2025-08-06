package base

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetErrorStatus(t *testing.T) {
	t.Run("returns 404 for gorm.ErrRecordNotFound", func(t *testing.T) {
		// Arrange
		err := gorm.ErrRecordNotFound

		// Act
		status := getErrorStatus(err)

		// Assert
		assert.Equal(t, fiber.StatusNotFound, status)
	})

	t.Run("returns CustomError code", func(t *testing.T) {
		// Arrange
		customErr := NewCustomError(fiber.StatusBadRequest, "validation failed", nil)

		// Act
		status := getErrorStatus(customErr)

		// Assert
		assert.Equal(t, fiber.StatusBadRequest, status)
	})

	t.Run("returns fiber.Error code", func(t *testing.T) {
		// Arrange
		fiberErr := &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
		}

		// Act
		status := getErrorStatus(fiberErr)

		// Assert
		assert.Equal(t, fiber.StatusUnauthorized, status)
	})

	t.Run("returns 500 for generic errors", func(t *testing.T) {
		// Arrange
		genericErr := errors.New("generic error")

		// Act
		status := getErrorStatus(genericErr)

		// Assert
		assert.Equal(t, fiber.StatusInternalServerError, status)
	})
}

func TestGetErrorData(t *testing.T) {
	t.Run("returns CustomError data", func(t *testing.T) {
		// Arrange
		data := CustomErrorData{"field": "value", "code": 123}
		customErr := NewCustomError(400, "test", data)

		// Act
		result := getErrorData(customErr)

		// Assert
		assert.Equal(t, data, result)
	})

	t.Run("returns nil for fiber.Error", func(t *testing.T) {
		// Arrange
		fiberErr := &fiber.Error{
			Code:    400,
			Message: "bad request",
		}

		// Act
		result := getErrorData(fiberErr)

		// Assert
		assert.Nil(t, result)
	})

	t.Run("returns nil for generic errors", func(t *testing.T) {
		// Arrange
		genericErr := errors.New("generic error")

		// Act
		result := getErrorData(genericErr)

		// Assert
		assert.Nil(t, result)
	})

	t.Run("returns nil for gorm errors", func(t *testing.T) {
		// Arrange
		gormErr := gorm.ErrRecordNotFound

		// Act
		result := getErrorData(gormErr)

		// Assert
		assert.Nil(t, result)
	})
}

func TestErrorHandler(t *testing.T) {
	t.Run("handles CustomError correctly", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			data := CustomErrorData{"field": "invalid", "reason": "validation_failed"}
			return NewCustomError(fiber.StatusBadRequest, "validation error", data)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response errorResponse
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.False(t, response.Success)
		assert.Equal(t, "validation error", response.Message)
		assert.NotNil(t, response.Data)
		assert.Equal(t, "invalid", response.Data["field"])
		assert.Equal(t, "validation_failed", response.Data["reason"])
	})

	t.Run("handles fiber.Error correctly", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			return &fiber.Error{
				Code:    fiber.StatusUnauthorized,
				Message: "access denied",
			}
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response errorResponse
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.False(t, response.Success)
		assert.Equal(t, "access denied", response.Message)
		assert.Nil(t, response.Data)
	})

	t.Run("handles gorm.ErrRecordNotFound correctly", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			return gorm.ErrRecordNotFound
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response errorResponse
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.False(t, response.Success)
		assert.Equal(t, "record not found", response.Message)
		assert.Nil(t, response.Data)
	})

	t.Run("handles generic errors correctly", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			return errors.New("something went wrong")
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response errorResponse
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.False(t, response.Success)
		assert.Equal(t, "something went wrong", response.Message)
		assert.Nil(t, response.Data)
	})

	t.Run("handles CustomError with nil data", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			return NewCustomError(fiber.StatusBadRequest, "error without data", nil)
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response errorResponse
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.False(t, response.Success)
		assert.Equal(t, "error without data", response.Message)
		// Data should be omitted from JSON when nil (omitempty tag)
		assert.Nil(t, response.Data)
	})

	t.Run("handles CustomError with empty data", func(t *testing.T) {
		// Arrange
		app := fiber.New(fiber.Config{
			ErrorHandler: ErrorHandler,
		})
		app.Get("/test", func(c *fiber.Ctx) error {
			return NewCustomError(fiber.StatusBadRequest, "error with empty data", CustomErrorData{})
		})

		req := httptest.NewRequest("GET", "/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response errorResponse
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.False(t, response.Success)
		assert.Equal(t, "error with empty data", response.Message)
		// Empty map should be serialized as {} and unmarshaled as non-nil empty map
		if response.Data != nil {
			assert.Len(t, response.Data, 0)
		}
	})
}
