package dalle

// ImageGenerationRequest is the request body for the image generation endpoint.
type ImageGenerationRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}
