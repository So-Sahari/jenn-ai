// Package ollama contains all ollama related calls
package ollama

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"jenn-ai/internal/fuzzy"

	"github.com/ollama/ollama/api"
)

func SelectOllamaModel(ctx context.Context) (string, error) {
	models, err := ListModels(ctx)
	if err != nil {
		return "", err
	}

	selectModel, err := selectModelInput(models)
	if err != nil {
		return "", err
	}
	return selectModel.ModelID, nil
}

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

func selectModelInput(models []OllamaModel) (OllamaModel, error) {
	var modelInfo OllamaModel

	selector := fuzzy.Prompter{}
	sortedModels := sortModelInputs(models)

	var modelsToSelect []string
	linePrefix := "#"

	for i, info := range sortedModels {
		models := fmt.Sprintf("Name: %s | ModelID: %s ", info.Name, info.ModelID)
		modelsToSelect = append(modelsToSelect, linePrefix+strconv.Itoa(i)+" "+models)
	}

	label := "Select your model"
	indexChoice, _ := selector.Select(label, modelsToSelect, fuzzy.FuzzySearchWithPrefixAnchor(modelsToSelect, linePrefix))

	fmt.Println()

	modelInfo = sortedModels[indexChoice]

	fmt.Printf("Selected model: %s - %s", modelInfo.Name, modelInfo.ModelID)
	fmt.Println()
	return modelInfo, nil
}

func sortModelInputs(models []OllamaModel) []OllamaModel {
	var sortedModels []OllamaModel

	sortedModels = append(sortedModels, models...)
	sort.Slice(sortedModels, func(i, j int) bool {
		return sortedModels[i].Name < sortedModels[j].Name
	})
	return sortedModels
}
