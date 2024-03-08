package cmd

import (
	"context"
	"fmt"

	"jenn-ai/app"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server for web",
	Long:  `Run server for web.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		initializeApp()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func initializeApp() {
	ctx := context.Background()
	mc := app.NewModelConfig(modelSource, model, region, temperature, topP, topK, maxTokens)
	mc.Serve(ctx)
}
