// Package models contains structs on model requests/responses
package models

// CommandModelInput contains the request
type CommandModelInput struct {
	Prompt            string   `json:"prompt"`             // The input text that serves as the starting point for generating the response.
	Temperature       float64  `json:"temperature"`        // Use a lower value to decrease randomness in the response.
	TopP              float64  `json:"p"`                  // Use a lower value to ignore less probable options. Set to 0 or 1.0 to disable.
	TopK              int      `json:"k"`                  // Specify the number of token choices the model uses to generate the next token.
	MaxTokensToSample int      `json:"max_tokens"`         // Specify the maximum number of tokens to use in the generated response.
	StopSequences     []string `json:"stop_sequences"`     // Configure up to four sequences that the model recognizes.
	ReturnLiklihoods  string   `json:"return_likelihoods"` // Specify how and if the token likelihoods are returned with the response.
	Stream            bool     `json:"stream"`             // Specify true to return the response piece-by-piece in real-time and false to return the complete response after the process finishes.
	NumGenerations    int      `json:"num_generations"`    // The maximum number of generations that the model should return.
}

// CommandModelOutput contains the response
type CommandModelOutput struct {
	Generations []CommandGeneration `json:"generations"` // A list of generated results along with the likelihoods for tokens requested.
}

// CommandGeneration contains the chunks of text
type CommandGeneration struct {
	ID   string `json:"id"`   // An identifier for the generation.
	Text string `json:"text"` // The generated text.
}
