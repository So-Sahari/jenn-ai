// Package models contains structs on model requests/responses
package models

// ClaudeModelInputs is needed for the request
type ClaudeModelInputs struct {
	Prompt            string  `json:"prompt"`
	MaxTokensToSample int     `json:"max_tokens_to_sample"`
	Temperature       float64 `json:"temperature,omitempty"`
	TopP              float64 `json:"top_p,omitempty"`
	TopK              int     `json:"top_k,omitempty"`
}

// ClaudeModelOutputs contains the response
type ClaudeModelOutputs struct {
	Completion string `json:"completion,omitempty"`
	StopReason string `json:"stop_reason,omitempty"`
	Stop       string `json:"stop,omitempty"`
}
