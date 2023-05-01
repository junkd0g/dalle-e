package dalle_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/junkd0g/dalle-e"
	"github.com/stretchr/testify/assert"
)

func TestNewClientWithInvalidAPIKey(t *testing.T) {

	t.Run("testing new client errors and success", func(t *testing.T) {
		tests := []struct {
			name     string
			apiKey   string
			domain   string
			expected error
		}{
			{
				name:     "empty API key",
				apiKey:   "",
				domain:   "",
				expected: errors.New("api key is required"),
			},
			{
				name:     "success",
				apiKey:   "sonme-api-key",
				domain:   "",
				expected: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := dalle.NewClient(tt.apiKey, tt.domain)
				assert.Equal(t, tt.expected, err)
			})
		}
	})

	t.Run("testing new client with options", func(t *testing.T) {
		apiKey := "test-api-key"
		domain := "https://example.com"
		timeout := time.Second * 30
		httpClient := &http.Client{Timeout: timeout}

		client, err := dalle.NewClient(
			apiKey,
			domain,
			dalle.WithHTTPClient(httpClient),
			dalle.WithTimeout(timeout),
		)
		assert.NoError(t, err)

		assert.Equal(t, apiKey, client.APIKey)
		assert.Equal(t, domain, client.Domain)
	})

	t.Run("testing domain", func(t *testing.T) {
		client, err := dalle.NewClient("test-api-key", "")
		assert.NoError(t, err)
		assert.Equal(t, dalle.DefaultDomain, client.Domain)

		client, err = dalle.NewClient("test-api-key", "https://example.com")
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com", client.Domain)
	})
}
