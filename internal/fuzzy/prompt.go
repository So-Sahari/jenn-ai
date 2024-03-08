// Package fuzzy contains logic to fuzzy search
package fuzzy

import (
	"log"

	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
)

// Input is used to interface with promptui
type Input interface {
	Select(label string, toSelect []string, searcher func(input string, index int) bool) (index int, value string)
	Enter(label string, dfault string) string
}

// Prompter is used to interface promptui
type Prompter struct{}

// Select is used to select from the prompt
func (p Prompter) Select(label string, toSelect []string, searcher func(input string, index int) bool) (int, string) {
	prompt := promptui.Select{
		Label:             label,
		Items:             toSelect,
		Size:              20,
		Searcher:          searcher,
		StartInSearchMode: true,
		Stdout:            SilentStdout,
	}
	index, value, err := prompt.Run()
	if err != nil {
		log.Fatalf("Error in prompt: %q", err)
	}
	return index, value
}

// Enter is used to pop open a selectable prompt
func (p Prompter) Enter(label string, dfault string) string {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   dfault,
		AllowEdit: false,
	}
	val, err := prompt.Run()
	if err != nil {
		log.Fatalf("Error in prompt: %q", err)
	}
	return val
}

type silentStdout struct{}

// SilentStdout is needed because users with terminal sounds enabled hear alerts
var SilentStdout = &silentStdout{}

// Write ensures no sounds from the terminal in stdout
func (s *silentStdout) Write(b []byte) (int, error) {
	if len(b) == 1 && b[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(b)
}

// Close closes the stdout
func (s *silentStdout) Close() error {
	return readline.Stdout.Close()
}
