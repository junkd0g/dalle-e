package dalle_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
	// Start a local HTTP test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Check if the request headers are set correctly
		if authHeader := r.Header.Get(authorizationStr); authHeader != testAPIKey {
			t.Errorf("Expected Authorization header %s, got %s", testAPIKey, authHeader)
		}
		if contentType := r.Header.Get(contentTypeStr); contentType != contentTypeValueStr {
			t.Errorf("Expected Content-Type header %s, got %s", contentTypeValueStr, contentType)
		}

		// Read request body and check if it matches expected values
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer r.Body.Close()

		var imgReq dalle.ImageGenerationRequest
		err = json.Unmarshal(reqBody, &imgReq)
		if err != nil {
			t.Fatal(err)
		}

		if imgReq.Prompt != testPrompt {
			t.Errorf("Expected prompt %s, got %s", testPrompt, imgReq.Prompt)
		}
		if imgReq.N != testN {
			t.Errorf("Expected N %d, got %d", testN, imgReq.N)
		}
		if imgReq.Size != testSize {
			t.Errorf("Expected size %s, got %s", testSize, imgReq.Size)
		}

		// Respond with a dummy image data
		fmt.Fprint(w, "dummy-image-data")
	}))
	defer ts.Close()

	// Create a new client with the test server URL as the domain
	client, err := dalle.NewClient(testAPIKey, ts.URL)
	assert.NoError(t, err)

	// Call the GenerateImageV1 function
	resp, err := client.GenerateImageV1(testPrompt, testN, testSize)
	assert.NoError(t, err)

	// Check if the response matches the expected value
	expectedResp := []byte("dummy-image-data")
	if !strings.EqualFold(string(resp), string(expectedResp)) {
		t.Errorf("Expected response %s, got %s", string(expectedResp), string(resp))
	}
}
