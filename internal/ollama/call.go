// Package ollama contains all ollama related calls
package ollama

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// OllamaModelConfig contains all settings
type OllamaModelConfig struct {
	ModelID     string  // ModelID is the ID of the model to invoke
	Temperature float64 // Temperature is part of model settings
	TopP        float64 // TopP is part of model settings
	TopK        int     // TopK is part of model settings
	MaxTokens   int     // MaxTokens is part of model settings
}

// NewModel is a factory
func NewModel(model string, temp, topP float64, topK, maxTokens int) OllamaModelConfig {
	return OllamaModelConfig{
		ModelID:     model,
		Temperature: temp,
		TopP:        topP,
		TopK:        topK,
		MaxTokens:   maxTokens,
	}
}

func (m *OllamaModelConfig) CallModel(ctx context.Context, message string) (string, error) {
	llm, err := ollama.New(ollama.WithModel(m.ModelID))
	if err != nil {
		return "", err
	}
	completion, err := llm.Call(ctx, message,
		llms.WithTemperature(m.Temperature),
		llms.WithTopP(m.TopP),
		llms.WithTopK(m.TopK),
		llms.WithMaxTokens(m.MaxTokens),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		return "", err
	}
	fmt.Println()

	// return fmt.Sprintf("%q", completion), nil
	return completion, nil
}
