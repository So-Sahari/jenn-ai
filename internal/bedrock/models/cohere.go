// Package models contains structs on model requests/responses
package models

// CommandModelInput contains the request
type CommandModelInput struct {
	Prompt            string   `json:"prompt"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"p"`
	TopK              int      `json:"k"`
	MaxTokensToSample int      `json:"max_tokens"`
	StopSequences     []string `json:"stop_sequences"`
	ReturnLiklihoods  string   `json:"return_likelihoods"`
	Stream            bool     `json:"stream"`
	NumGenerations    int      `json:"num_generations"`
}

// CommandModelOutput contains the response
type CommandModelOutput struct {
	Generations []CommandGeneration `json:"generations"`
}

// CommandGeneration contains the chunks of text
type CommandGeneration struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
