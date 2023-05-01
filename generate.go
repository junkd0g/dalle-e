package dalle

import (
	"encoding/json"
	"fmt"
)

// GenerateImageV1 generates an image from a prompt.
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
