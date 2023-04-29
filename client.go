package dalle

import (
	"net/http"
	"time"
)

const (
	// DefaultDomain is the default domain to use for the API.
	DefaultDomain = "https://api.openai.com"
)

// Client is the main struct for the Dalle API client.
type Client struct {
	// The API key to use for authentication.
	APIKey string
	// The Domain to use for the API.
	Domain string
	// The HTTP client to use for requests.
	HTTPClient http.Client
}

// NewClient creates a new Dalle API client.
func NewClient(apiKey string, domain string, opts ...ClientOption) (*Client, error) {
	client := &Client{
		APIKey:     apiKey,
		Domain:     domain,
		HTTPClient: http.Client{},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// ClientOption is a function that modifies the client's configuration.
type ClientOption func(*Client)

// WithHTTPClient sets the HTTP client to use.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = *httpClient
	}
}

// WithTimeout sets the timeout for API requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.HTTPClient.Timeout = timeout
	}
}
