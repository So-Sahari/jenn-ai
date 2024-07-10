// Package app
package app

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var completion string

type ModelConfig struct {
	Platform string
	ModelID  string

	Temperature float64
	TopP        float64
	TopK        int
	MaxTokens   int

	Region string // used for AWS models
}

func NewModelConfig(platform, modelID, region string, temp, topP float64, topK, maxTokens int) ModelConfig {
	return ModelConfig{
		Platform:    platform,
		ModelID:     modelID,
		Temperature: temp,
		TopP:        topP,
		TopK:        topK,
		MaxTokens:   maxTokens,
		Region:      region,
	}
}

func (mc *ModelConfig) Serve(ctx context.Context) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	initDB()
	defer db.Close()

	router.GET("/", renderIndex)
	router.GET("/current-state", getCurrentState)
	router.GET("/messages", getAllMessagesFromDB)
	router.GET("/message/:id/render", getMessagesFromDB)
	router.GET("/model-platform", getModelPlatform)
	router.GET("/model", getModelsByPlatform(ctx, mc.Region))
	router.POST("/select-model", selectModel)
	router.POST("/run", mc.runModel(ctx))
	router.POST("/new-conversation", createConversation)

	if err := router.Run(":31000"); err != nil {
		log.Fatal(err)
	}
}

func renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": "JennAI",
	})
}
