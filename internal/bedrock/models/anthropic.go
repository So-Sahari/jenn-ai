// Package models contains structs on model requests/responses
package models

// ClaudeModelInputs is needed to marshal the legacy request type
// Supported Models: claude-instant-v1, claude-v1, claude-v2, claude-v2:1
type ClaudeModelInputs struct {
	Prompt            string  `json:"prompt"`                // The prompt that you want Claude to complete. For proper response generation you need to format your prompt using alternating \n\nHuman: and \n\nAssistant: conversational turns.
	MaxTokensToSample int     `json:"max_tokens_to_sample"`  // The maximum number of tokens to generate before stopping.
	Temperature       float64 `json:"temperature,omitempty"` // The amount of randomness injected into the response.
	TopP              float64 `json:"top_p,omitempty"`       // Use nucleus sampling.
	TopK              int     `json:"top_k,omitempty"`       // Only sample from the top K options for each subsequent token.
}

// ClaudeModelOutputs is needed to unmarshal the legacy response
// Supported Models: claude-instant-v1, claude-v1, claude-v2, claude-v2:1
type ClaudeModelOutputs struct {
	Completion string `json:"completion,omitempty"`  // The resulting completion up to and excluding the stop sequences.
	StopReason string `json:"stop_reason,omitempty"` // The reason why the model stopped generating the response.
	Stop       string `json:"stop,omitempty"`        // contains the stop sequence that signalled the model to stop generating text.
}

// ClaudeMessagesInput is needed to marshal the new request type
// Supported Models: claude-instant-v1.2, claude-v2, claude-v2.1, claude-v3
type ClaudeMessagesInput struct {
	AnthropicVersion string          `json:"anthropic_version"`        // The anthropic version. The value must be bedrock-2023-05-31.
	Messages         []ClaudeMessage `json:"messages"`                 // The input messages.
	MaxTokens        int             `json:"max_tokens"`               // The maximum number of tokens to generate before stopping.
	Temperature      float64         `json:"temperature"`              // The amount of randomness injected into the response.
	TopK             int             `json:"top_k"`                    // Only sample from the top K options for each subsequent token.
	TopP             float64         `json:"top_p"`                    // Use nucleus sampling.
	System           string          `json:"system,omitempty"`         // The system prompt for the request. Which provides context, instructions, and guidelines to Claude before presenting it with a question or task.
	StopSequences    []string        `json:"stop_sequences,omitempty"` // The stop sequences that signal the model to stop generating text.
}

// ClaudeMessage contains messages
type ClaudeMessage struct {
	Role    string          `json:"role"`    // The role of the conversation turn. Valid values are user and assistant.
	Content []ClaudeContent `json:"content"` // The content of the conversation turn.
}

// ClaudeContent contains content
type ClaudeContent struct {
	Type string `json:"type"`           // The type of the content. Valid values are image and text.
	Text string `json:"text,omitempty"` // The text content.
}

// ClaudeMessagesOutput is needed to unmarshal the new request type
// Supported Models: claude-instant-v1.2, claude-v2, claude-v2.1, claude-v3
type ClaudeMessagesOutput struct {
	Type  string    `json:"type"`  // The type of the response. Valid values are image and text.
	Index int       `json:"index"` // The index of the response.
	Delta TextDelta `json:"delta"` // The delta of the response.
}

// TextDelta contains type and text
type TextDelta struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
