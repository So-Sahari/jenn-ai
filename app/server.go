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

	router.GET("/", renderIndex)
	router.GET("/model-platform", getModelPlatform)
	router.POST("/select-platform", setModelPlatform)
	router.GET("/model", getModelsByPlatform(ctx, mc.Region))
	router.POST("/select-model", selectModel)
	router.POST("/run", mc.runModel(ctx, mc.Region))

	if err := router.Run(":31000"); err != nil {
		log.Fatal(err)
	}
}

func renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": "JennAI",
	})
}
