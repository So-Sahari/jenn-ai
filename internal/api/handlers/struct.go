package handlers

import (
	"html/template"

	"jenn-ai/internal/config"
)

type ModelConfig config.ModelConfig

// NewParameters creates a new ModelConfig
func NewParameters(platform, modelID, region string, temp, topP float64, topK, maxTokens int) ModelConfig {
	return ModelConfig{
		Platform:    platform,
		ModelID:     modelID,
		Temperature: temp,
		TopP:        topP,
		TopK:        topK,
		MaxTokens:   maxTokens,
		Region:      region,
	}
}

type ChatMessage struct {
	Human    template.HTML
	Response template.HTML
	Platform string
	Model    string
}
