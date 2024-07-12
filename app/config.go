package app

type ModelConfig struct {
	Platform string
	ModelID  string

	Temperature float64
	TopP        float64
	TopK        int
	MaxTokens   int

	Region string // used for AWS models
}

func NewModelConfig(platform, modelID, region string, temp, topP float64, topK, maxTokens int) ModelConfig {
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
