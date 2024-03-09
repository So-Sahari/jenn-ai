// Package cmd contains all cobra commands
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	region      string
	model       string
	temperature float64
	topP        float64
	topK        int
	maxTokens   int
	modelSource string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "jenn-ai",
		Short: "call LLMs from the commandline and through a server",
		Long:  `call LLMs from the commandline and through a server`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "AWS region")
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "", "The model id")
	rootCmd.PersistentFlags().StringVarP(&modelSource, "model-source", "s", "", "The model source (e.g. Bedrock, Ollama)")

	rootCmd.PersistentFlags().Float64VarP(&temperature, "temperature", "t", 1, "temperature setting")
	rootCmd.PersistentFlags().Float64VarP(&topP, "topP", "", 0.999, "topP setting")
	rootCmd.PersistentFlags().IntVarP(&topK, "topK", "", 250, "topK setting")
	rootCmd.PersistentFlags().IntVarP(&maxTokens, "max-tokens", "", 500, "max tokens to sample")
}
