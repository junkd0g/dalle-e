package dalle

import (
	"errors"
	"net/http"
	"time"
)

const (
	// DefaultDomain is the default API endpoint used if no custom domain is provided.
	DefaultDomain = "https://api.openai.com"
)

// Client is the main struct for the Dalle API client.
type Client struct {
	// The API key to use for authentication.
	APIKey string
	// The Domain to use for the API.
	Domain string
	// The HTTP client to use for requests.
	Client http.Client
}

// NewClient initializes and returns a new instance of the Dalle API client.
// It requires a non-empty API key and optionally accepts a custom domain.
// Additional configuration can be applied using functional options.
func NewClient(apiKey string, domain string, opts ...ClientOption) (*Client, error) {

	if apiKey == "" {
		return nil, errors.New("api key is required")
	}

	if domain == "" {
		domain = DefaultDomain
	}

	client := &Client{
		APIKey: apiKey,
		Domain: domain,
		Client: http.Client{},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// ClientOption defines a function type for configuring the Client.
// It allows for modular and flexible client configuration.
type ClientOption func(*Client)

// WithHTTPClient returns a ClientOption that sets a custom HTTP client for making API requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.Client = *httpClient
	}
}

// WithTimeout returns a ClientOption that configures the timeout duration for API requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.Client.Timeout = timeout
	}
}
