//go:build !dev

package api

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestOptionsProvider(t *testing.T) {
	t.Run("returns production options", func(t *testing.T) {
		options := OptionsProvider()

		assert.NotNil(t, options)
		assert.Equal(t, "https", options.Protocol)
		assert.NotNil(t, options.Configure)
	})

	t.Run("configure function sets production settings", func(t *testing.T) {
		options := OptionsProvider()
		config := &fiber.Config{}

		options.Configure(config)

		assert.True(t, config.DisableStartupMessage)
	})

	t.Run("returns new instance on each call", func(t *testing.T) {
		options1 := OptionsProvider()
		options2 := OptionsProvider()

		assert.NotSame(t, options1, options2)
		// Both should have same configuration
		assert.Equal(t, options1.Protocol, options2.Protocol)
	})

	t.Run("production protocol is HTTPS", func(t *testing.T) {
		options := OptionsProvider()

		assert.Equal(t, "https", options.Protocol)
	})
}

func TestRunHandler(t *testing.T) {
	t.Run("transforms GET request correctly", func(t *testing.T) {
		// Create a simple HTTP handler for testing
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"hello"}`))
		})

		// Create a Lambda function URL request
		req := events.LambdaFunctionURLRequest{
			RawPath: "/api/test",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
			Headers: map[string]string{
				"Authorization": "Bearer token123",
				"User-Agent":    "test-agent",
			},
		}

		// Act
		response, err := runHandler(context.Background(), handler, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "application/json", response.Headers["Content-Type"])
		assert.Equal(t, `{"message":"hello"}`, response.Body)
	})

	t.Run("transforms POST request with body correctly", func(t *testing.T) {
		// Create a handler that echoes the request body
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"received":"` + string(body) + `"}`))
		})

		// Create a Lambda function URL request with body
		req := events.LambdaFunctionURLRequest{
			RawPath: "/api/users",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "POST",
				},
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"name":"John","email":"john@example.com"}`,
		}

		// Act
		response, err := runHandler(context.Background(), handler, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 201, response.StatusCode)
		assert.Equal(t, "application/json", response.Headers["Content-Type"])
		assert.Contains(t, response.Body, `"received":"`)
		assert.Contains(t, response.Body, `"name":"John"`)
		assert.Contains(t, response.Body, `"email":"john@example.com"`)
	})

	t.Run("handles query parameters correctly", func(t *testing.T) {
		// Create a handler that returns query parameters
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			param1 := r.URL.Query().Get("param1")
			param2 := r.URL.Query().Get("param2")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"param1":"` + param1 + `","param2":"` + param2 + `"}`))
		})

		// Create a Lambda function URL request with query parameters
		req := events.LambdaFunctionURLRequest{
			RawPath:        "/api/search",
			RawQueryString: "param1=value1&param2=value2",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
		}

		// Act
		response, err := runHandler(context.Background(), handler, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, response.StatusCode)
		assert.Contains(t, response.Body, `"param1":"value1"`)
		assert.Contains(t, response.Body, `"param2":"value2"`)
	})

	t.Run("handles empty query string", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})

		req := events.LambdaFunctionURLRequest{
			RawPath:        "/api/test",
			RawQueryString: "", // Empty query string
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
		}

		response, err := runHandler(context.Background(), handler, req)

		assert.NoError(t, err)
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "ok", response.Body)
	})

	t.Run("forwards all headers correctly", func(t *testing.T) {
		var receivedHeaders http.Header
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedHeaders = r.Header
			w.WriteHeader(http.StatusOK)
		})

		req := events.LambdaFunctionURLRequest{
			RawPath: "/api/test",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
			Headers: map[string]string{
				"Authorization": "Bearer token123",
				"Content-Type":  "application/json",
				"X-Custom":      "custom-value",
				"User-Agent":    "test-client/1.0",
			},
		}

		_, err := runHandler(context.Background(), handler, req)

		assert.NoError(t, err)
		assert.Equal(t, "Bearer token123", receivedHeaders.Get("Authorization"))
		assert.Equal(t, "application/json", receivedHeaders.Get("Content-Type"))
		assert.Equal(t, "custom-value", receivedHeaders.Get("X-Custom"))
		assert.Equal(t, "test-client/1.0", receivedHeaders.Get("User-Agent"))
	})

	t.Run("handles multiple response headers correctly", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Add("Set-Cookie", "session=abc123")
			w.Header().Add("Set-Cookie", "theme=dark")
			w.Header().Set("X-Rate-Limit", "1000")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"ok"}`))
		})

		req := events.LambdaFunctionURLRequest{
			RawPath: "/api/test",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
		}

		response, err := runHandler(context.Background(), handler, req)

		assert.NoError(t, err)
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "application/json", response.Headers["Content-Type"])
		assert.Equal(t, "session=abc123, theme=dark", response.Headers["Set-Cookie"])
		assert.Equal(t, "1000", response.Headers["X-Rate-Limit"])
		assert.Equal(t, `{"status":"ok"}`, response.Body)
	})

	t.Run("handles error responses correctly", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"not found"}`))
		})

		req := events.LambdaFunctionURLRequest{
			RawPath: "/api/nonexistent",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
		}

		response, err := runHandler(context.Background(), handler, req)

		assert.NoError(t, err)
		assert.Equal(t, 404, response.StatusCode)
		assert.Equal(t, "application/json", response.Headers["Content-Type"])
		assert.Equal(t, `{"error":"not found"}`, response.Body)
	})

	t.Run("sets RequestURI correctly", func(t *testing.T) {
		var receivedRequestURI string
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedRequestURI = r.RequestURI
			w.WriteHeader(http.StatusOK)
		})

		req := events.LambdaFunctionURLRequest{
			RawPath:        "/api/test",
			RawQueryString: "param=value",
			RequestContext: events.LambdaFunctionURLRequestContext{
				HTTP: events.LambdaFunctionURLRequestContextHTTPDescription{
					Method: "GET",
				},
			},
		}

		_, err := runHandler(context.Background(), handler, req)

		assert.NoError(t, err)
		assert.Equal(t, "/api/test?param=value", receivedRequestURI)
	})
}
