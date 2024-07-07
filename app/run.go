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
	"jenn-ai/internal/parser"
	"jenn-ai/internal/state"

	"github.com/gin-gonic/gin"
)

func (mc *ModelConfig) runModel(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		appState := state.GetState()
		platform := appState.GetPlatform()
		modelID := appState.GetModel()
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

			parsed = strings.ReplaceAll(parsed, "<pre>", "<div class='card bg-base-100 shadow-xl'><div class='card-body'><pre>")
			parsed = strings.ReplaceAll(parsed, "</pre>", "</pre></div></div>")
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"Human":    template.HTML(message),
				"Response": template.HTML(parsed),
				"Platform": platform,
				"Model":    modelID,
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

			parsed = strings.ReplaceAll(parsed, "<pre>", "<div class='card bg-base-100 shadow-xl'><div class='card-body'><pre>")
			parsed = strings.ReplaceAll(parsed, "</pre>", "</pre></div></div>")
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"Human":    template.HTML(message),
				"Response": template.HTML(parsed),
				"Platform": platform,
				"Model":    modelID,
			})
		default:
			fmt.Println("No Model Platform selected or unsupported")
		}
	}
}
