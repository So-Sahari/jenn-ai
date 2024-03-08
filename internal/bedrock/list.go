// Package bedrock contains aws logic
package bedrock

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"jenn-ai/internal/fuzzy"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
)

// FoundationModel contains model fields
type FoundationModel struct {
	Name     string
	Provider string
	ID       string
	Modality string
}

// ClientAPI is used to interface with bedrock client
type ClientAPI interface {
	ListFoundationModels(ctx context.Context, params *bedrock.ListFoundationModelsInput, optFns ...func(*bedrock.Options)) (*bedrock.ListFoundationModelsOutput, error)
}

// SelectBedrockModel is used to list and search Bedrock Models
func SelectBedrockModel(ctx context.Context, client *bedrock.Client) (string, error) {
	fm, err := ListModels(ctx, client)
	if err != nil {
		return "", err
	}
	selectedModel := selectModelInput(fm)
	return selectedModel.ID, nil
}

func ListModels(ctx context.Context, api ClientAPI) ([]FoundationModel, error) {
	var output []FoundationModel

	response, err := api.ListFoundationModels(ctx, &bedrock.ListFoundationModelsInput{})
	if err != nil {
		return output, err
	}

	for _, model := range response.ModelSummaries {
		output = append(output, FoundationModel{
			Name:     *model.ModelName,
			Provider: *model.ProviderName,
			ID:       *model.ModelId,
			Modality: fmt.Sprintf("%v", model.OutputModalities),
		})
	}
	return output, nil
}

func selectModelInput(models []FoundationModel) FoundationModel {
	selector := fuzzy.Prompter{}
	sortedModels := sortModelInputs(models)

	var modelsToSelect []string
	linePrefix := "#"

	for i, info := range sortedModels {
		modelInfo := fmt.Sprintf("Name: %s | Provider: %s | Id: %s | Modality: %s", info.Name, info.Provider, info.ID, info.Modality)
		modelsToSelect = append(modelsToSelect, linePrefix+strconv.Itoa(i)+" "+modelInfo)
	}

	label := "Select your model"
	indexChoice, _ := selector.Select(label, modelsToSelect, fuzzy.FuzzySearchWithPrefixAnchor(modelsToSelect, linePrefix))

	fmt.Println()

	modelInfo := sortedModels[indexChoice]

	fmt.Printf("Selected model: %s - %s", modelInfo.Name, modelInfo.ID)
	fmt.Println()
	return modelInfo
}

func sortModelInputs(models []FoundationModel) []FoundationModel {
	var sortedModels []FoundationModel

	sortedModels = append(sortedModels, models...)
	sort.Slice(sortedModels, func(i, j int) bool {
		return sortedModels[i].Name < sortedModels[j].Name
	})
	return sortedModels
}
