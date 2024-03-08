// Package ollama contains all ollama related calls
package ollama

// OllamaModel contains the base struct for search and sort
type OllamaModel struct {
	Name    string
	ModelID string
}

// OllamaModels contains popular and more likely used models
var OllamaModels = []OllamaModel{
	{
		Name:    "Llama 2",
		ModelID: "llama2",
	},
	{
		Name:    "Mistral",
		ModelID: "mistral",
	},
	{
		Name:    "Mixtral",
		ModelID: "mixtral",
	},
	{
		Name:    "Dolphin Mixtral",
		ModelID: "dolphin-mixtral",
	},
	{
		Name:    "Dolphin Phi",
		ModelID: "dolphin-phi",
	},
	{
		Name:    "Phi-2",
		ModelID: "phi",
	},
	{
		Name:    "Neural Chat",
		ModelID: "neural-chat",
	},
	{
		Name:    "Starling",
		ModelID: "starling-lm",
	},
	{
		Name:    "Code Llama",
		ModelID: "codellama",
	},
	{
		Name:    "Llama 2 Uncensored",
		ModelID: "llama2-uncensored",
	},
	{
		Name:    "Orca Mini",
		ModelID: "orca-mini",
	},
	{
		Name:    "Vicuna",
		ModelID: "vicuna",
	},
	{
		Name:    "LLaVA",
		ModelID: "llava",
	},
	{
		Name:    "Gemma 2B",
		ModelID: "gemma:2b",
	},
	{
		Name:    "Gemma 7B",
		ModelID: "gemma:7b",
	},
}
