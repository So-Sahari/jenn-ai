package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/So-Sahari/go-bedrock"
	"github.com/gin-gonic/gin"

	"jenn-ai/internal/ollama"
	"jenn-ai/internal/state"
)

func getModelsByPlatform(ctx context.Context, region string) gin.HandlerFunc {
	return func(c *gin.Context) {
		platform := c.Query("platform-option")
		appState := state.GetState()
		appState.SetPlatform(platform)

		var models string
		var modelList []string

		switch platform {
		case "Bedrock":
			client, err := bedrock.CreateBedrockClient(ctx, region)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Bedrock client"})
			}

			bm, err := bedrock.ListModels(ctx, client)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to list models"})
			}
			for _, v := range bm {
				modelList = append(modelList, v.ID)
			}
		case "Ollama":
			ollamaModels, err := ollama.ListModels(ctx)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to list models"})
			}
			for _, n := range ollamaModels {
				modelList = append(modelList, n.ModelID)
			}
		default:
			fmt.Println("No Model Platform selected")
		}

		models = "<option disabled selected>Select Model</option>"
		for _, m := range modelList {
			models += fmt.Sprintf("<option>%s</option>", m)
		}
		c.Data(http.StatusOK, "text/html", []byte(models))
	}
}

func selectModel(c *gin.Context) {
	model := c.PostForm("model-option")
	appState := state.GetState()
	appState.SetModel(model)
}
