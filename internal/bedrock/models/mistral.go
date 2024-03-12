// Package models contains structs on model requests/responses
package models

// MistralRequest contains the request needed for mistral models
type MistralRequest struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	TopK        int     `json:"top_k,omitempty"`
}

// MistralResponse contains the response obtained from mistral models
type MistralResponse struct {
	Outputs []MistralOutput `json:"outputs"`
}

// MistralOutput contains the response text and stop response
type MistralOutput struct {
	Text         string `json:"text"`
	StopResponse string `json:"stop_response"`
}
