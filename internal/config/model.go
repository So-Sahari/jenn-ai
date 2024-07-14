package config

import "jenn-ai/internal/db"

// ModelConfig contains the parameters for the model
type ModelConfig struct {
	Platform    string
	ModelID     string
	Temperature float64
	TopP        float64
	TopK        int
	MaxTokens   int

	Region string // used for AWS models

	DB db.Client // client for db
}
