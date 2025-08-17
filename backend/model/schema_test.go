package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchema_ScreenshotPath(t *testing.T) {
	t.Run("returns nil when ScreenshotedAt is nil", func(t *testing.T) {
		// Arrange
		schema := &Schema{
			Identifiable:   &Identifiable{ID: MustStringToID("550e8400-e29b-41d4-a716-446655440000")},
			ScreenshotedAt: nil,
		}

		// Act
		result := schema.ScreenshotPath()

		// Assert
		assert.Nil(t, result)
	})

	t.Run("returns path with version when ScreenshotedAt is set", func(t *testing.T) {
		// Arrange
		testTime := time.Unix(1234567890, 0)
		testID := MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		schema := &Schema{
			Identifiable:   &Identifiable{ID: testID},
			ScreenshotedAt: &testTime,
		}

		// Act
		result := schema.ScreenshotPath()

		// Assert
		assert.NotNil(t, result)
		expectedPath := "data/images/550e8400-e29b-41d4-a716-446655440000/schema.webp?v=1234567890"
		assert.Equal(t, expectedPath, *result)
	})

	t.Run("returns correct path for different timestamp", func(t *testing.T) {
		// Arrange
		testTime := time.Unix(9876543210, 0)
		testID := MustStringToID("123e4567-e89b-12d3-a456-426614174000")
		schema := &Schema{
			Identifiable:   &Identifiable{ID: testID},
			ScreenshotedAt: &testTime,
		}

		// Act
		result := schema.ScreenshotPath()

		// Assert
		assert.NotNil(t, result)
		expectedPath := "data/images/123e4567-e89b-12d3-a456-426614174000/schema.webp?v=9876543210"
		assert.Equal(t, expectedPath, *result)
	})
}

func TestSchemaScreenshotKey(t *testing.T) {
	t.Run("returns correct key format", func(t *testing.T) {
		// Arrange
		testID := MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		// Act
		result := SchemaScreenshotKey(testID)

		// Assert
		expected := "data/images/550e8400-e29b-41d4-a716-446655440000/schema.webp"
		assert.Equal(t, expected, result)
	})

	t.Run("returns correct key for different UUID", func(t *testing.T) {
		// Arrange
		testID := MustStringToID("123e4567-e89b-12d3-a456-426614174000")

		// Act
		result := SchemaScreenshotKey(testID)

		// Assert
		expected := "data/images/123e4567-e89b-12d3-a456-426614174000/schema.webp"
		assert.Equal(t, expected, result)
	})

	t.Run("handles zero UUID", func(t *testing.T) {
		// Arrange
		testID := MustStringToID("00000000-0000-0000-0000-000000000000")

		// Act
		result := SchemaScreenshotKey(testID)

		// Assert
		expected := "data/images/00000000-0000-0000-0000-000000000000/schema.webp"
		assert.Equal(t, expected, result)
	})
}
