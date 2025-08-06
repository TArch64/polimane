package base

import (
	"maps"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomError(t *testing.T) {
	t.Run("creates custom error with code, message and data", func(t *testing.T) {
		// Arrange
		code := 400
		message := "test error"
		data := CustomErrorData{"field": "value"}

		// Act
		err := NewCustomError(code, message, data)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, code, err.Code)
		assert.Equal(t, message, err.Message)
		assert.Equal(t, data, err.Data)
	})

	t.Run("creates custom error with nil data", func(t *testing.T) {
		// Arrange
		code := 500
		message := "internal error"

		// Act
		err := NewCustomError(code, message, nil)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, code, err.Code)
		assert.Equal(t, message, err.Message)
		assert.Nil(t, err.Data)
	})

	t.Run("creates custom error with empty data", func(t *testing.T) {
		// Arrange
		code := 404
		message := "not found"
		data := CustomErrorData{}

		// Act
		err := NewCustomError(code, message, data)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, code, err.Code)
		assert.Equal(t, message, err.Message)
		assert.Equal(t, data, err.Data)
	})
}

func TestCustomError_Error(t *testing.T) {
	t.Run("returns error message", func(t *testing.T) {
		// Arrange
		message := "test error message"
		err := NewCustomError(400, message, nil)

		// Act
		result := err.Error()

		// Assert
		assert.Equal(t, message, result)
	})

	t.Run("implements error interface", func(t *testing.T) {
		// Arrange
		err := NewCustomError(400, "test", nil)

		// Act & Assert
		var e error = err
		assert.NotNil(t, e)
		assert.Equal(t, "test", e.Error())
	})
}

func TestCustomError_AddCustomData(t *testing.T) {
	t.Run("adds single custom data", func(t *testing.T) {
		// Arrange
		originalData := CustomErrorData{"field1": "value1"}
		err := NewCustomError(400, "test", originalData)
		additionalData := CustomErrorData{"field2": "value2"}

		// Act
		newErr := err.AddCustomData(additionalData)

		// Assert
		assert.NotNil(t, newErr)
		assert.NotEqual(t, err, newErr) // Should be a new instance
		assert.Equal(t, err.Code, newErr.Code)
		assert.Equal(t, err.Message, newErr.Message)
		assert.Equal(t, "value1", newErr.Data["field1"])
		assert.Equal(t, "value2", newErr.Data["field2"])
		// Original error should remain unchanged
		assert.NotContains(t, err.Data, "field2")
	})

	t.Run("adds multiple custom data", func(t *testing.T) {
		// Arrange
		originalData := CustomErrorData{"field1": "value1"}
		err := NewCustomError(400, "test", originalData)
		data1 := CustomErrorData{"field2": "value2"}
		data2 := CustomErrorData{"field3": "value3"}

		// Act
		newErr := err.AddCustomData(data1, data2)

		// Assert
		assert.NotNil(t, newErr)
		assert.Equal(t, "value1", newErr.Data["field1"])
		assert.Equal(t, "value2", newErr.Data["field2"])
		assert.Equal(t, "value3", newErr.Data["field3"])
	})

	t.Run("overwrites duplicate keys with later values", func(t *testing.T) {
		// Arrange
		originalData := CustomErrorData{"field": "original"}
		err := NewCustomError(400, "test", originalData)
		data1 := CustomErrorData{"field": "first"}
		data2 := CustomErrorData{"field": "second"}

		// Act
		newErr := err.AddCustomData(data1, data2)

		// Assert
		assert.Equal(t, "second", newErr.Data["field"])
	})

	t.Run("handles nil original data", func(t *testing.T) {
		// Arrange
		err := NewCustomError(400, "test", nil)
		additionalData := CustomErrorData{"field": "value"}

		// Act
		newErr := err.AddCustomData(additionalData)

		// Assert
		assert.NotNil(t, newErr.Data)
		assert.Equal(t, "value", newErr.Data["field"])
	})

	t.Run("handles empty additional data", func(t *testing.T) {
		// Arrange
		originalData := CustomErrorData{"field": "value"}
		err := NewCustomError(400, "test", originalData)

		// Act
		newErr := err.AddCustomData()

		// Assert
		assert.NotNil(t, newErr)
		assert.Equal(t, originalData, newErr.Data)
	})

	t.Run("handles nil additional data", func(t *testing.T) {
		// Arrange
		originalData := CustomErrorData{"field": "value"}
		err := NewCustomError(400, "test", originalData)

		// Act
		newErr := err.AddCustomData(nil)

		// Assert
		assert.NotNil(t, newErr)
		assert.Equal(t, originalData, newErr.Data)
	})
}

func TestNewReasonedError(t *testing.T) {
	t.Run("creates custom error with reason in data", func(t *testing.T) {
		// Arrange
		code := 404
		reason := "ResourceNotFound"

		// Act
		err := NewReasonedError(code, reason)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, code, err.Code)
		assert.Equal(t, reason, err.Message)
		assert.NotNil(t, err.Data)
		assert.Equal(t, reason, err.Data["reason"])
	})

	t.Run("creates error with empty reason", func(t *testing.T) {
		// Arrange
		code := 400
		reason := ""

		// Act
		err := NewReasonedError(code, reason)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, code, err.Code)
		assert.Equal(t, reason, err.Message)
		assert.Equal(t, reason, err.Data["reason"])
	})
}

func TestNotFoundErr(t *testing.T) {
	t.Run("is pre-defined NotFound error", func(t *testing.T) {
		// Act & Assert
		assert.NotNil(t, NotFoundErr)
		assert.Equal(t, fiber.StatusNotFound, NotFoundErr.Code)
		assert.Equal(t, "NotFound", NotFoundErr.Message)
		assert.Equal(t, "NotFound", NotFoundErr.Data["reason"])
	})

	t.Run("implements error interface", func(t *testing.T) {
		// Act & Assert
		var err error = NotFoundErr
		assert.Equal(t, "NotFound", err.Error())
	})
}

func TestCustomErrorData(t *testing.T) {
	t.Run("can be used as map", func(t *testing.T) {
		// Arrange
		data := CustomErrorData{
			"field1": "value1",
			"field2": 123,
			"field3": true,
		}

		// Act & Assert
		assert.Equal(t, "value1", data["field1"])
		assert.Equal(t, 123, data["field2"])
		assert.Equal(t, true, data["field3"])
	})

	t.Run("can be copied with maps.Copy", func(t *testing.T) {
		// Arrange
		source := CustomErrorData{"field1": "value1"}
		dest := CustomErrorData{}

		// Act
		maps.Copy(dest, source)

		// Assert
		assert.Equal(t, "value1", dest["field1"])
	})
}
