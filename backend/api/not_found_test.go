package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestApiNotFound(t *testing.T) {
	t.Run("returns 404 status code", func(t *testing.T) {
		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request
		req := httptest.NewRequest("GET", "/nonexistent", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
	})

	t.Run("returns JSON error response", func(t *testing.T) {
		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request
		req := httptest.NewRequest("GET", "/api/invalid", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		// Check Content-Type header
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		// Check response body
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Not Found", response["error"])
	})

	t.Run("logs the unhandled route path", func(t *testing.T) {
		// Capture log output
		var logOutput bytes.Buffer
		oldOutput := log.Writer()
		log.SetOutput(&logOutput)
		defer log.SetOutput(oldOutput)

		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request
		req := httptest.NewRequest("GET", "/test/path/not/found", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		// Check that the path was logged
		logString := logOutput.String()
		assert.Contains(t, logString, "Unhandled route:")
		assert.Contains(t, logString, "/test/path/not/found")
	})

	t.Run("handles root path", func(t *testing.T) {
		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request for root path
		req := httptest.NewRequest("GET", "/", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		// Check response body
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Not Found", response["error"])
	})

	t.Run("handles paths with query parameters", func(t *testing.T) {
		// Capture log output to verify path logging
		var logOutput bytes.Buffer
		oldOutput := log.Writer()
		log.SetOutput(&logOutput)
		defer log.SetOutput(oldOutput)

		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request with query parameters
		req := httptest.NewRequest("GET", "/api/search?q=test&limit=10", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		// Verify the path (without query params) is logged
		logString := logOutput.String()
		assert.Contains(t, logString, "Unhandled route:")
		assert.Contains(t, logString, "/api/search")
	})

	t.Run("handles different HTTP methods", func(t *testing.T) {
		methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

		for _, method := range methods {
			t.Run("method_"+method, func(t *testing.T) {
				// Create a Fiber app for testing
				app := fiber.New()

				// Register the not found handler
				app.Use(apiNotFound)

				// Create a test request with specific method
				req := httptest.NewRequest(method, "/api/endpoint", nil)

				// Perform the request
				resp, err := app.Test(req)

				// Assert
				assert.NoError(t, err)
				assert.Equal(t, 404, resp.StatusCode)

				// Check response body
				var response map[string]interface{}
				err = json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, "Not Found", response["error"])
			})
		}
	})

	t.Run("response format is valid JSON", func(t *testing.T) {
		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request
		req := httptest.NewRequest("GET", "/invalid", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		// Verify it's valid JSON with expected structure
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Not Found", response["error"])

		// Verify only expected fields are present
		assert.Len(t, response, 1)
	})

	t.Run("does not log sensitive information", func(t *testing.T) {
		// Capture log output
		var logOutput bytes.Buffer
		oldOutput := log.Writer()
		log.SetOutput(&logOutput)
		defer log.SetOutput(oldOutput)

		// Create a Fiber app for testing
		app := fiber.New()

		// Register the not found handler
		app.Use(apiNotFound)

		// Create a test request with path that might contain sensitive info
		req := httptest.NewRequest("GET", "/api/users/password-reset?token=secret123", nil)

		// Perform the request
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		// Check log only contains the path, not query parameters with sensitive data
		logString := logOutput.String()
		assert.Contains(t, logString, "Unhandled route:")
		assert.Contains(t, logString, "/api/users/password-reset")
		// The log should not contain query parameters
		assert.NotContains(t, logString, "secret123")
	})
}
