package workos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/workos/workos-go/v4/pkg/workos_errors"
)

func TestGetErrorCode(t *testing.T) {
	tests := []struct {
		name         string
		httpError    workos_errors.HTTPError
		expectedCode string
	}{
		{
			name: "valid error code in response body",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": "email_verification_code_expired", "message": "Email verification code has expired"}`,
			},
			expectedCode: CodeEmailVerificationCodeExpired,
		},
		{
			name: "custom error code in response body",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": "custom_error_code", "message": "Custom error occurred"}`,
			},
			expectedCode: "custom_error_code",
		},
		{
			name: "empty code in response body",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": "", "message": "Error without code"}`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "missing code field in response body",
			httpError: workos_errors.HTTPError{
				RawBody: `{"message": "Error without code field"}`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "invalid JSON in response body",
			httpError: workos_errors.HTTPError{
				RawBody: `invalid json response`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "empty response body",
			httpError: workos_errors.HTTPError{
				RawBody: "",
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "malformed JSON - missing quotes",
			httpError: workos_errors.HTTPError{
				RawBody: `{code: email_verification_code_expired}`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "malformed JSON - incomplete",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": "email_verification_code_expired"`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "null code value",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": null, "message": "Error with null code"}`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "numeric code value",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": 123, "message": "Error with numeric code"}`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "boolean code value",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": true, "message": "Error with boolean code"}`,
			},
			expectedCode: CodeUnknownError,
		},
		{
			name: "nested JSON structure",
			httpError: workos_errors.HTTPError{
				RawBody: `{"error": {"code": "nested_error_code"}, "message": "Nested error"}`,
			},
			expectedCode: CodeUnknownError, // Because the code is nested, not at root level
		},
		{
			name: "extra fields with valid code",
			httpError: workos_errors.HTTPError{
				RawBody: `{"code": "valid_code", "message": "Error message", "timestamp": "2023-01-01T00:00:00Z", "details": {"field": "value"}}`,
			},
			expectedCode: "valid_code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetErrorCode(tt.httpError)
			assert.Equal(t, tt.expectedCode, result)
		})
	}
}

func TestErrorCodeConstants(t *testing.T) {
	// Test that our constants have the expected values
	assert.Equal(t, "email_verification_code_expired", CodeEmailVerificationCodeExpired)
	assert.Equal(t, "unknown", CodeUnknownError)
}
