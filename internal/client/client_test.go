package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Get(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		body := "test response"
		baseURL := "http://test.com"
		path := "/test"
		ctx := context.Background()

		// Create mock transport that returns a successful response with the test body.
		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(body))),
			},
		}

		c := New(baseURL, &http.Client{Transport: transport})

		resp, err := c.Get(ctx, baseURL+path)

		assert.NoError(t, err)
		assert.Equal(t, []byte(body), resp)
	})

	t.Run("nil context error", func(t *testing.T) {
		baseURL := "http://test.com"
		path := "/notfound"

		// Create mock transport that returns a not found error.
		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(bytes.NewReader(nil)),
			},
		}

		c := New(baseURL, &http.Client{Transport: transport})

		_, err := c.Get(nil, baseURL+path)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "nil Context")
	})

	t.Run("request with error", func(t *testing.T) {
		baseURL := "http://test.com"
		path := "/notfound"
		ctx := context.Background()

		// Create mock transport that returns a not found error.
		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(bytes.NewReader(nil)),
			},
		}

		c := New(baseURL, &http.Client{Transport: transport})

		_, err := c.Get(ctx, baseURL+path)

		assert.Error(t, err)
	})

	t.Run("unsuccessful response", func(t *testing.T) {
		baseURL := "http://test.com"
		path := "/test"
		ctx := context.Background()

		// Create mock transport that returns a internal server error.
		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewReader(nil)),
			},
		}

		c := New(baseURL, &http.Client{Transport: transport})

		_, err := c.Get(ctx, baseURL+path)

		assert.Error(t, err)
	})

	t.Run("timeout", func(t *testing.T) {
		// Create a custom HTTP server with a delay handler that waits for longer than the timeout period.
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Second)
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		// Create an HTTP client with custom transport that uses the custom server.
		tr := &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		}
		client := &http.Client{
			Timeout:   1 * time.Second,
			Transport: tr,
		}

		// Create a new client with the custom HTTP client.
		c := New(ts.URL, client)

		// Call the Get method and expect an error due to the request timing out.
		ctx := context.Background()
		_, err := c.Get(ctx, "test")

		assert.Error(t, err)
		urlError, ok := err.(*url.Error)
		require.True(t, ok)
		assert.ErrorContains(t, urlError, context.DeadlineExceeded.Error())
	})

	t.Run("body read error", func(t *testing.T) {
		// Create a custom HTTP server with a delay handler that waits for longer than the timeout period.
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Second)
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		// Create a mock response with an unreadable body.
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       errorReader{},
		}

		// Create an HTTP client with custom transport that uses the custom server.
		tr := &mockTransport{
			response: resp,
		}
		client := &http.Client{
			Timeout:   1 * time.Second,
			Transport: tr,
		}

		// Create a new client with the custom HTTP client.
		c := New(ts.URL, client)

		// Call the Get method and expect an error due to the request timing out.
		ctx := context.Background()
		_, err := c.Get(ctx, "test")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "readall error")
	})
}

// errorReader is a custom io.Reader that always returns an error when Read is called.
type errorReader struct{}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("readall error")
}

func (e errorReader) Close() error {
	return nil
}

// mockTransport is a mock http.RoundTripper that delays the response and returns a specified response.
type mockTransport struct {
	delay    time.Duration
	response *http.Response
}

// RoundTrip delays the response for the specified duration and returns the specified response.
func (t *mockTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	time.Sleep(t.delay)
	return t.response, nil
}
