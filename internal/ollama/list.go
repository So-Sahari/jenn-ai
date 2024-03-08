// Package ollama contains all ollama related calls
package ollama

import (
	"fmt"
	"sort"
	"strconv"

	"jenn-ai/internal/fuzzy"
)

func SelectOllamaModel() string {
	return selectModelInput().ModelID
}

func selectModelInput() OllamaModel {
	models := OllamaModels
	selector := fuzzy.Prompter{}
	sortedModels := sortModelInputs(models)

	var modelsToSelect []string
	linePrefix := "#"

	for i, info := range sortedModels {
		modelInfo := fmt.Sprintf("Name: %s | ModelID: %s ", info.Name, info.ModelID)
		modelsToSelect = append(modelsToSelect, linePrefix+strconv.Itoa(i)+" "+modelInfo)
	}

	label := "Select your model"
	indexChoice, _ := selector.Select(label, modelsToSelect, fuzzy.FuzzySearchWithPrefixAnchor(modelsToSelect, linePrefix))

	fmt.Println()

	modelInfo := sortedModels[indexChoice]

	fmt.Printf("Selected model: %s - %s", modelInfo.Name, modelInfo.ModelID)
	fmt.Println()
	return modelInfo
}

func sortModelInputs(models []OllamaModel) []OllamaModel {
	var sortedModels []OllamaModel

	sortedModels = append(sortedModels, models...)
	sort.Slice(sortedModels, func(i, j int) bool {
		return sortedModels[i].Name < sortedModels[j].Name
	})
	return sortedModels
}
