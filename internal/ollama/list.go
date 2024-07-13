// Package ollama contains all ollama related calls
package ollama

import (
	"context"

	"github.com/ollama/ollama/api"
)

func ListModels(ctx context.Context) ([]OllamaModel, error) {
	var models []OllamaModel

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return models, err
	}

	localModels, err := client.List(ctx)
	if err != nil {
		return models, err
	}

	for _, model := range localModels.Models {
		m := OllamaModel{Name: model.Name, ModelID: model.Model}
		models = append(models, m)
	}
	return models, nil
}
