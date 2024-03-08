// Package fuzzy contains logic to fuzzy search
package fuzzy

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// FuzzySearchWithPrefixAnchor is used to fuzzy search through a generic string list
func FuzzySearchWithPrefixAnchor(itemsToSelect []string, linePrefix string) func(input string, index int) bool {
	return func(input string, index int) bool {
		role := itemsToSelect[index]

		if strings.HasPrefix(input, linePrefix) {
			return strings.HasPrefix(role, input)
		}

		if fuzzy.MatchFold(input, role) {
			return true
		}
		return false
	}
}

// ModelSource are the supported sources to run models
var ModelSource = []string{
	"Bedrock",
	"Ollama",
}

// GetModelSource is used to setup a prompt for region
func GetModelSource(prompt Input) string {
	_, source := prompt.Select("Select your model source: ", ModelSource, func(input string, index int) bool {
		target := ModelSource[index]
		return fuzzy.MatchFold(input, target)
	})
	return source
}
