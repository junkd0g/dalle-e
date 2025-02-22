package dalle

import (
	"encoding/json"
	"fmt"
)

// GenerateImageV1 sends a request to generate an image from a provided prompt.
// It takes a prompt string, the number of images (n), and the size as parameters.
// The function returns the response as a byte slice or an error if any.
func (c *Client) GenerateImageV1(prompt string, n int, size string) ([]byte, error) {
	reqBody, err := json.Marshal(ImageGenerationRequest{
		Prompt: prompt,
		N:      n,
		Size:   size,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	resp, err := c.post("/images/generations", "/v1", string(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	return resp, nil
}
