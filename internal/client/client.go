package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// HttpClient interface that specifies the behavior of an HTTP client.
type HttpClient interface {
	Get(ctx context.Context, suffix string) ([]byte, error)
}

// Client struct that implements the HttpClient interface.
type Client struct {
	baseURL string
	c       *http.Client
}

// New returns a new Client instance with the given base URL and HTTP client
func New(baseURL string, c *http.Client) *Client {
	return &Client{
		baseURL: baseURL,
		c:       c,
	}
}

// Get makes a GET request to the API with the given suffix and returns the response body as a []byte
func (c *Client) Get(ctx context.Context, suffix string) ([]byte, error) {
	// Construct the full URL for the API endpoint.
	uri := fmt.Sprintf("%s/%s", c.baseURL, suffix)

	// Create a new HTTP request with the given context, method, URL, and body.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err // If an error occurred while creating the request, return it
	}

	// Enable HTTP keep-alive to improve performance.
	req.Close = true

	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", res.Status)
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
