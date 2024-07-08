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
		conversationID := appState.GetConversationID()
		message := c.PostForm("prompt")

		var response string
		var completion strings.Builder
		var err error

		if conversationID == 0 {
			conversationID, err = createNewConversation()
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new conversation"})
				return
			}
			appState.SetConversationID(conversationID)
		} else {
			// Retrieve previous messages in the conversation to build the completion context
			previousMessages, err := getMessagesByConversationID(conversationID)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			for _, msg := range previousMessages {
				completion.WriteString(msg.Human + " " + msg.Response)
			}
		}
		completion.WriteString(message)

		switch platform {
		case "Bedrock":
			brClient, err := bedrock.CreateBedrockruntimeClient(ctx, mc.Region)
			if err != nil {
				log.Fatalf("encountered error with client: %v", err)
			}
			model := bedrock.NewModel(modelID, mc.Temperature, mc.TopP, mc.TopK, mc.MaxTokens)

			response, err = model.InvokeModel(ctx, brClient, completion.String())
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

		case "Ollama":
			model := ollama.NewModel(modelID, mc.Temperature, mc.TopP, mc.TopK, mc.MaxTokens)
			response, err = model.CallModel(ctx, completion.String())
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

		default:
			fmt.Println("No Model Platform selected or unsupported")
		}

		// Insert message into the database
		if err := insertMessage(conversationID, message, response, platform, modelID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// parse markdown
		parsed, err := parser.ParseMD(response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		parsed = strings.ReplaceAll(parsed, "<pre>", "<div class='card bg-base-100 shadow-xl'><div class='card-body'><pre>")
		parsed = strings.ReplaceAll(parsed, "</pre>", "</pre></div></div>")
		c.HTML(http.StatusOK, "chat.html", gin.H{
			"Human":          template.HTML(message),
			"Response":       template.HTML(parsed),
			"Platform":       platform,
			"Model":          modelID,
			"ConversationID": conversationID,
		})
	}
}
