package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"jenn-ai/internal/api/handlers"
	"jenn-ai/internal/config"
	"jenn-ai/internal/db"
	"jenn-ai/internal/state"
)

func Serve(ctx context.Context, c config.ModelConfig) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	mc := handlers.NewParameters(c.Platform, c.ModelID, c.Region, c.Temperature, c.TopP, c.TopK, c.MaxTokens)

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

	router.GET("/", handlers.RenderIndex)
	router.GET("/state", handlers.GetCurrentState)
	router.GET("/messages", mc.GetAllMessages())
	router.GET("/message/:id/render", mc.GetMessage())
	router.GET("/model-platform", handlers.GetModelPlatform)
	router.GET("/model", handlers.GetModelsByPlatform(ctx, mc.Region))
	router.GET("/parameters", handlers.GetParameterState)

	router.POST("/select-model", handlers.SelectModel)
	router.POST("/run", mc.Invoke(ctx))
	router.POST("/parameters", handlers.SetParameterState)
	router.POST("/new-conversation", mc.CreateConversation())

	router.DELETE("/conversation/:id/delete", mc.DeleteChat())

	if err := router.Run(":31000"); err != nil {
		log.Fatal(err)
	}
}