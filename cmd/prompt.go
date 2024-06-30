// Package cmd contains all cobra commands
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	b "jenn-ai/internal/bedrock"
	"jenn-ai/internal/fuzzy"
	o "jenn-ai/internal/ollama"

	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Write a prompt to send to models",
	Long: `Write a prompt to send to models.
  Currently Supports:
  - Bedrock Foundational models 
  - Ollama (local) models`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		if modelSource == "" {
			prompt := fuzzy.Prompter{}
			modelSource = fuzzy.GetModelSource(prompt)
		}

		switch modelSource {
		case "Bedrock":
			invokeBedrockModel(ctx)
		case "Ollama":
			invokeOllamaModel(ctx)
		default:
			log.Fatal("Unable to determine model source")
		}
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}

func invokeOllamaModel(ctx context.Context) {
	if model == "" {
		var err error
		model, err = o.SelectOllamaModel(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

	var chatMessages string
	reader := bufio.NewReader(os.Stdin)
	model := o.NewModel(model, temperature, topP, topK, maxTokens)

	fmt.Println("Enter a prompt and then press enter:")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		message := chatMessages + input

		response, err := model.CallModel(ctx, message)
		if err != nil {
			log.Fatal(err)
		}
		chatMessages = message + response
	}
}

func invokeBedrockModel(ctx context.Context) {
	bClient, err := b.CreateBedrockClient(ctx, region)
	if err != nil {
		log.Fatalf("encountered error with client: %v", err)
	}
	brClient, err := b.CreateBedrockruntimeClient(ctx, region)
	if err != nil {
		log.Fatalf("encountered error with client: %v", err)
	}

	if model == "" {
		model, err = b.SelectBedrockModel(ctx, bClient)
		if err != nil {
			log.Fatalf("encountered error selecting bedrock model: %v", err)
		}
	}

	var chatMessages string
	reader := bufio.NewReader(os.Stdin)
	model := b.NewModel(model, temperature, topP, topK, maxTokens)

	fmt.Println("Enter a prompt and then press enter:")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		message := chatMessages + input

		response, err := model.InvokeModel(ctx, brClient, message)
		if err != nil {
			log.Fatal(err)
		}
		chatMessages = message + response
	}
}
