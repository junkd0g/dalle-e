package dalle_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/junkd0g/dalle-e"
	"github.com/stretchr/testify/assert"
)

const (
	authorizationStr    = "Authorization"
	contentTypeStr      = "Content-Type"
	contentTypeValueStr = "application/json"
	testAPIKey          = "test-api-key"
	testPrompt          = "Test Prompt"
	testN               = 2
	testSize            = "256x256"
)

func TestGenerateImageV1(t *testing.T) {
	var receivedRequest *http.Request
	var requestBody []byte

	// Start a test HTTP server that returns dummy image data.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the incoming request for further assertions.
		receivedRequest = r

		var err error
		requestBody, err = io.ReadAll(r.Body)
		assert.NoError(t, err, "Should be able to read request body")
		defer r.Body.Close()

		// Validate that the request method is POST.
		assert.Equal(t, http.MethodPost, r.Method, "HTTP method should be POST")

		// Write dummy image data to the response body.
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("dummy-image-data"))
		assert.NoError(t, err, "Should be able to write response body")
	}))
	defer ts.Close()

	// Create a new Dalle client using the test server URL as the domain.
	client, err := dalle.NewClient(testAPIKey, ts.URL, dalle.WithTimeout(5*time.Second))
	assert.NoError(t, err, "NewClient should not return an error")

	// Call the GenerateImageV1 function to generate an image.
	resp, err := client.GenerateImageV1(testPrompt, testN, testSize)
	assert.NoError(t, err, "GenerateImageV1 should not return an error")
	assert.Equal(t, []byte("dummy-image-data"), resp, "Response data should match expected dummy data")

	// Validate the request headers.
	assert.NotNil(t, receivedRequest, "A request should have been received by the test server")
	expectedAuthHeader := "Bearer " + testAPIKey
	assert.Equal(t, expectedAuthHeader, receivedRequest.Header.Get(authorizationStr), "Authorization header should include the 'Bearer' prefix")
	assert.Equal(t, contentTypeValueStr, receivedRequest.Header.Get(contentTypeStr), "Content-Type header should be application/json")

	// Verify the JSON request body.
	var imgReq dalle.ImageGenerationRequest
	err = json.Unmarshal(requestBody, &imgReq)
	assert.NoError(t, err, "Request body should be valid JSON")
	assert.Equal(t, testPrompt, imgReq.Prompt, "Prompt should match expected value")
	assert.Equal(t, testN, imgReq.N, "Number of images (n) should match expected value")
	assert.Equal(t, testSize, imgReq.Size, "Size should match expected value")
}
