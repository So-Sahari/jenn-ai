package app

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"jenn-ai/internal/bedrock"
	"jenn-ai/internal/ollama"
	"jenn-ai/internal/parser"

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

			// parse markdown
			parsed, err := parser.ParseMD(response)
			if err != nil {
				log.Fatal(err)
			}

			c.HTML(http.StatusOK, "chat.html", gin.H{
				"Human":    template.HTML("<b>You:</b><br>" + message),
				"Response": template.HTML("<b>JennAI:</b><br>" + parsed),
			})
		case "Ollama":
			model := ollama.NewModel(modelID, mc.Temperature, mc.TopP, mc.TopK, mc.MaxTokens)
			input := completion + message
			response, err := model.CallModel(ctx, input)
			if err != nil {
				log.Fatal(err)
			}
			completion = input + response

			// parse markdown
			parsed, err := parser.ParseMD(response)
			if err != nil {
				log.Fatal(err)
			}

			c.HTML(http.StatusOK, "chat.html", gin.H{
				"Human":    template.HTML("<b>You:</b><br>" + message),
				"Response": template.HTML("<b>JennAI:</b><br>" + parsed),
			})
		default:
			fmt.Println("No Model Platform selected or unsupported")
		}
	}
}
