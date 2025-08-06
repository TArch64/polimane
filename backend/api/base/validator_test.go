package base

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type testValidationStruct struct {
	Name     string `validate:"required,min=2,max=50"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"min=0,max=120"`
	Password string `validate:"required,min=8"`
}

type testValidationOptional struct {
	Name  string `validate:"omitempty,min=2"`
	Email string `validate:"omitempty,email"`
}

func TestInitValidator(t *testing.T) {
	t.Run("initializes validator instance", func(t *testing.T) {
		// Arrange
		validatorInstance = nil

		// Act
		InitValidator()

		// Assert
		assert.NotNil(t, validatorInstance)
		assert.NotNil(t, validatorInstance.validator)
		assert.IsType(t, &validator.Validate{}, validatorInstance.validator)
	})

	t.Run("can be called multiple times safely", func(t *testing.T) {
		// Act
		InitValidator()
		firstInstance := validatorInstance
		InitValidator()
		secondInstance := validatorInstance

		// Assert
		assert.NotNil(t, firstInstance)
		assert.NotNil(t, secondInstance)
		// Note: We don't check if they're the same instance since InitValidator
		// creates a new instance each time, which is the current behavior
	})
}

func TestRequestValidator_Validate(t *testing.T) {
	// Ensure validator is initialized
	InitValidator()

	t.Run("returns no errors for valid data", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "John Doe",
			Email:    "john@example.com",
			Age:      25,
			Password: "password123",
		}

		// Act
		errors := validatorInstance.Validate(data)

		// Assert
		assert.Empty(t, errors)
	})

	t.Run("returns validation errors for invalid data", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "", // Required field missing
			Email:    "invalid-email",
			Age:      -5,      // Below minimum
			Password: "short", // Too short
		}

		// Act
		errors := validatorInstance.Validate(data)

		// Assert
		assert.NotEmpty(t, errors)
		assert.Len(t, errors, 4) // Should have 4 validation errors

		// Check that all errors have Error set to true
		for _, err := range errors {
			assert.True(t, err.Error)
		}

		// Check field names are present
		fieldNames := make([]string, len(errors))
		for i, err := range errors {
			fieldNames[i] = err.FailedField
		}
		assert.Contains(t, fieldNames, "Name")
		assert.Contains(t, fieldNames, "Email")
		assert.Contains(t, fieldNames, "Age")
		assert.Contains(t, fieldNames, "Password")
	})

	t.Run("returns specific validation error details", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "A", // Too short (min=2)
			Email:    "john@example.com",
			Age:      25,
			Password: "password123",
		}

		// Act
		errors := validatorInstance.Validate(data)

		// Assert
		assert.Len(t, errors, 1)
		assert.Equal(t, "Name", errors[0].FailedField)
		assert.Equal(t, "min", errors[0].Tag)
		assert.Equal(t, "A", errors[0].Value)
		assert.True(t, errors[0].Error)
	})

	t.Run("handles struct with no validation tags", func(t *testing.T) {
		// Arrange
		type noValidationStruct struct {
			Name string
			Age  int
		}
		data := noValidationStruct{
			Name: "",
			Age:  -1,
		}

		// Act
		errors := validatorInstance.Validate(data)

		// Assert
		assert.Empty(t, errors)
	})

	t.Run("handles optional fields correctly", func(t *testing.T) {
		// Arrange
		data := testValidationOptional{
			Name:  "", // Empty but optional
			Email: "", // Empty but optional
		}

		// Act
		errors := validatorInstance.Validate(data)

		// Assert
		assert.Empty(t, errors)
	})

	t.Run("validates optional fields when provided", func(t *testing.T) {
		// Arrange
		data := testValidationOptional{
			Name:  "A",             // Too short (min=2)
			Email: "invalid-email", // Invalid email
		}

		// Act
		errors := validatorInstance.Validate(data)

		// Assert
		assert.Len(t, errors, 2)
		fieldNames := make([]string, len(errors))
		for i, err := range errors {
			fieldNames[i] = err.FailedField
		}
		assert.Contains(t, fieldNames, "Name")
		assert.Contains(t, fieldNames, "Email")
	})
}

func TestValidate(t *testing.T) {
	// Ensure validator is initialized
	InitValidator()

	t.Run("returns nil for valid data", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "John Doe",
			Email:    "john@example.com",
			Age:      25,
			Password: "password123",
		}

		// Act
		err := Validate(data)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("returns fiber.Error for invalid data", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "", // Required field missing
			Email:    "john@example.com",
			Age:      25,
			Password: "password123",
		}

		// Act
		err := Validate(data)

		// Assert
		assert.Error(t, err)
		fiberErr, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusBadRequest, fiberErr.Code)
		assert.Contains(t, fiberErr.Message, "Name")
		assert.Contains(t, fiberErr.Message, "required")
	})

	t.Run("includes multiple validation errors in message", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "",              // Required field missing
			Email:    "invalid-email", // Invalid email
			Age:      25,
			Password: "password123",
		}

		// Act
		err := Validate(data)

		// Assert
		assert.Error(t, err)
		fiberErr, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Contains(t, fiberErr.Message, "Name")
		assert.Contains(t, fiberErr.Message, "Email")
		assert.Contains(t, fiberErr.Message, " and ") // Multiple errors joined by " and "
	})

	t.Run("formats validation error message correctly", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "A", // Too short (min=2)
			Email:    "john@example.com",
			Age:      25,
			Password: "password123",
		}

		// Act
		err := Validate(data)

		// Assert
		assert.Error(t, err)
		fiberErr, ok := err.(*fiber.Error)
		assert.True(t, ok)

		// Check that the message follows the format: [FieldName]: 'Value' | Needs to implement 'Tag'
		assert.Contains(t, fiberErr.Message, "[Name]")
		assert.Contains(t, fiberErr.Message, "'A'")
		assert.Contains(t, fiberErr.Message, "Needs to implement 'min'")
	})

	t.Run("handles edge case with age boundary", func(t *testing.T) {
		// Arrange
		data := testValidationStruct{
			Name:     "John Doe",
			Email:    "john@example.com",
			Age:      121, // Above maximum (max=120)
			Password: "password123",
		}

		// Act
		err := Validate(data)

		// Assert
		assert.Error(t, err)
		fiberErr, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Contains(t, fiberErr.Message, "Age")
		assert.Contains(t, fiberErr.Message, "'121'")
		assert.Contains(t, fiberErr.Message, "max")
	})

	t.Run("returns nil when no validation errors", func(t *testing.T) {
		// Arrange
		type noValidationStruct struct {
			Name string
		}
		data := noValidationStruct{Name: "anything"}

		// Act
		err := Validate(data)

		// Assert
		assert.NoError(t, err)
	})
}

func TestValidationErrorResponse(t *testing.T) {
	t.Run("has correct structure", func(t *testing.T) {
		// Arrange
		response := ValidationErrorResponse{
			Error:       true,
			FailedField: "Name",
			Tag:         "required",
			Value:       "",
		}

		// Act & Assert
		assert.True(t, response.Error)
		assert.Equal(t, "Name", response.FailedField)
		assert.Equal(t, "required", response.Tag)
		assert.Equal(t, "", response.Value)
	})

	t.Run("can hold different value types", func(t *testing.T) {
		// Test with string value
		stringResponse := ValidationErrorResponse{
			Error:       true,
			FailedField: "Name",
			Tag:         "min",
			Value:       "short",
		}
		assert.Equal(t, "short", stringResponse.Value)

		// Test with integer value
		intResponse := ValidationErrorResponse{
			Error:       true,
			FailedField: "Age",
			Tag:         "min",
			Value:       -5,
		}
		assert.Equal(t, -5, intResponse.Value)
	})
}
