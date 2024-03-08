package app

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"jenn-ai/internal/bedrock"
	"jenn-ai/internal/ollama"

	"github.com/gin-gonic/gin"
)

func (mc *ModelConfig) runModel(ctx context.Context, region string) gin.HandlerFunc {
	return func(c *gin.Context) {
		message := c.PostForm("prompt")
		switch platform {
		case "Bedrock":
			brClient, err := bedrock.CreateBedrockruntimeClient(ctx, mc.Region)
			if err != nil {
				log.Fatalf("encountered error with client: %v", err)
			}
			model := bedrock.NewModel(modelID, mc.Temperature, mc.TopP, mc.TopK, mc.MaxTokens)

			input := completion + message
			response, err := model.InvokeModel(ctx, brClient, input)
			if err != nil {
				log.Fatal(err)
			}
			completion = input + response

			processPrompt := strings.ReplaceAll(message, "\n", "<br>")
			processPrompt = strings.ReplaceAll(processPrompt, "    ", `<span class="tab-space"></span>`)
			processResponse := strings.ReplaceAll(response, "\n", "<br>")
			processResponse = strings.ReplaceAll(processResponse, "    ", `<span class="tab-space"></span>`)
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"Human":    template.HTML("<b>You:</b><br>" + processPrompt),
				"Response": template.HTML("<b>JennAI:</b><br>" + processResponse),
			})
		case "Ollama":
			model := ollama.NewModel(modelID, mc.Temperature, mc.TopP, mc.TopK, mc.MaxTokens)
			input := completion + message
			response, err := model.CallModel(ctx, input)
			if err != nil {
				log.Fatal(err)
			}
			completion = input + response

			processPrompt := strings.ReplaceAll(message, "\n", "<br>")
			processPrompt = strings.ReplaceAll(processPrompt, "    ", `<span class="tab-space"></span>`)
			processResponse := strings.ReplaceAll(response, "\n", "<br>")
			processResponse = strings.ReplaceAll(processResponse, "    ", `<span class="tab-space"></span>`)
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"Human":    template.HTML("<b>You:</b><br>" + processPrompt),
				"Response": template.HTML("<b>JennAI:</b><br>" + processResponse),
			})
		default:
			fmt.Println("No Model Platform selected or unsupported")
		}
	}
}
