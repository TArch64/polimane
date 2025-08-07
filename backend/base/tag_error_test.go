package base

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagError(t *testing.T) {
	t.Run("returns nil when error is nil", func(t *testing.T) {
		result := TagError("test.tag", nil)

		assert.Nil(t, result)
	})

	t.Run("wraps error with tag", func(t *testing.T) {
		originalErr := errors.New("original error")
		tag := "test.operation"

		result := TagError(tag, originalErr)

		assert.Error(t, result)
		assert.Equal(t, "test.operation: original error", result.Error())
		assert.ErrorIs(t, result, originalErr)
	})

	t.Run("works with empty tag", func(t *testing.T) {
		originalErr := errors.New("original error")

		result := TagError("", originalErr)

		assert.Error(t, result)
		assert.Equal(t, ": original error", result.Error())
		assert.ErrorIs(t, result, originalErr)
	})

	t.Run("works with empty error message", func(t *testing.T) {
		originalErr := errors.New("")
		tag := "test.tag"

		result := TagError(tag, originalErr)

		assert.Error(t, result)
		assert.Equal(t, "test.tag: ", result.Error())
		assert.ErrorIs(t, result, originalErr)
	})

	t.Run("preserves error wrapping chain", func(t *testing.T) {
		wrappedErr := errors.New("wrapped error")
		tag := "service.method"

		result := TagError(tag, wrappedErr)

		assert.Error(t, result)
		assert.Equal(t, "service.method: wrapped error", result.Error())
		assert.ErrorIs(t, result, wrappedErr)
		// Verify the original error is still accessible through unwrapping
		assert.Equal(t, wrappedErr, errors.Unwrap(result))
	})

	t.Run("handles complex tag with dots", func(t *testing.T) {
		originalErr := errors.New("database connection failed")
		tag := "user.repository.create"

		result := TagError(tag, originalErr)

		assert.Error(t, result)
		assert.Equal(t, "user.repository.create: database connection failed", result.Error())
		assert.ErrorIs(t, result, originalErr)
	})

	t.Run("handles tag with special characters", func(t *testing.T) {
		originalErr := errors.New("validation failed")
		tag := "api/v1/users#create"

		result := TagError(tag, originalErr)

		assert.Error(t, result)
		assert.Equal(t, "api/v1/users#create: validation failed", result.Error())
		assert.ErrorIs(t, result, originalErr)
	})

	t.Run("can be chained multiple times", func(t *testing.T) {
		originalErr := errors.New("network timeout")

		// First level tagging
		level1 := TagError("http.client", originalErr)
		// Second level tagging
		level2 := TagError("user.service", level1)

		assert.Error(t, level2)
		assert.Equal(t, "user.service: http.client: network timeout", level2.Error())
		assert.ErrorIs(t, level2, level1)
		assert.ErrorIs(t, level2, originalErr)
	})
}
