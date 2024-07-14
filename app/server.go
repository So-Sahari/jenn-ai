// Package app
package app

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"jenn-ai/internal/db"
	"jenn-ai/internal/state"
)

func (mc *ModelConfig) Serve(ctx context.Context) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	clientDB, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer clientDB.Conn.Close()
	mc.DB = clientDB

	// save model parameters to state
	state := state.GetState()
	state.SetMaxTokens(mc.MaxTokens)
	state.SetTemperature(mc.Temperature)
	state.SetTopP(mc.TopP)
	state.SetTopK(mc.TopK)

	router.GET("/", renderIndex)
	router.GET("/current-state", getCurrentState)
	router.GET("/messages", mc.getAllMessagesFromDB())
	router.GET("/message/:id/render", mc.getMessagesFromDB())
	router.GET("/model-platform", getModelPlatform)
	router.GET("/model", getModelsByPlatform(ctx, mc.Region))
	router.GET("/parameters", getParameterState)
	router.POST("/select-model", selectModel)
	router.POST("/run", mc.runModel(ctx))
	router.POST("/parameters", setParameterState)
	router.POST("/new-conversation", mc.createConversation())
	router.DELETE("/conversation/:id/delete", mc.deleteChat())

	if err := router.Run(":31000"); err != nil {
		log.Fatal(err)
	}
}

func renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "JennAI",
	})
}
