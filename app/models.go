package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"jenn-ai/internal/bedrock"
	"jenn-ai/internal/ollama"

	"github.com/gin-gonic/gin"
)

var modelID string

func getModelsByPlatform(ctx context.Context, region string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var models string
		var modelList []string
		switch platform {
		case "Bedrock":
			client, err := bedrock.CreateBedrockClient(ctx, region)
			if err != nil {
				log.Fatal("Unable to create Bedrock client")
			}

			bm, err := bedrock.ListModels(ctx, client)
			if err != nil {
				log.Fatal("Unable to list models")
			}
			for _, v := range bm {
				modelList = append(modelList, v.ID)
			}
		case "Ollama":
			for _, n := range ollama.OllamaModels {
				modelList = append(modelList, n.ModelID)
			}
		default:
			fmt.Println("No Model Platform selected")
		}
		for _, m := range modelList {
			models += fmt.Sprintf(`<a href="#" 
      data-value="%s"
      hx-post="/select-model?option=%s" 
      class="model-item block px-4 py-2 text-sm text-black-700 hover:bg-blue-400">
      %s</a>`, m, m, m)
		}
		c.Data(http.StatusOK, "text/html", []byte(models))
	}
}

func selectModel(c *gin.Context) {
	selectedOption := c.Query("option")
	fmt.Printf("Received the following: %s\n", selectedOption)
	modelID = selectedOption
}
