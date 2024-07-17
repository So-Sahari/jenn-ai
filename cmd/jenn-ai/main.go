package main

import (
	"context"
	"flag"

	"jenn-ai/internal/api"
	"jenn-ai/internal/config"
)

var (
	platform    = flag.String("platform", "", "The Platform (e.g. Bedrock, Ollama)")
	model       = flag.String("model", "", "The model id")
	region      = flag.String("region", "us-east-1", "AWS region")
	temperature = flag.Float64("temperature", 1, "temperature setting")
	topP        = flag.Float64("topP", 0.999, "topP setting")
	topK        = flag.Int("topK", 250, "topK setting")
	maxTokens   = flag.Int("max-tokens", 500, "max tokens to sample")
	dbPath      = flag.String("db-path", "./chat.db", "The path to the sqlite database")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	mc := config.ModelConfig{
		Platform:    *platform,
		ModelID:     *model,
		Region:      *region,
		Temperature: *temperature,
		TopP:        *topP,
		TopK:        *topK,
		MaxTokens:   *maxTokens,
	}
	api.Serve(ctx, mc, *dbPath)
}
