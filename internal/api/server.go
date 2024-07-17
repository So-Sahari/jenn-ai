package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"jenn-ai/internal/api/handlers"
	"jenn-ai/internal/config"
	"jenn-ai/internal/db"
	"jenn-ai/internal/fs"
	"jenn-ai/internal/state"
)

func Serve(ctx context.Context, c config.ModelConfig, dbPath string) {
	router := gin.Default()

	// Use the embedded template set
	tmpl := fs.GetTemplates()
	router.SetHTMLTemplate(tmpl)

	mc := handlers.NewParameters(c.Platform, c.ModelID, c.Region, c.Temperature, c.TopP, c.TopK, c.MaxTokens)

	err := db.NewDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// save model parameters to state
	state := state.GetState()
	state.SetMaxTokens(mc.MaxTokens)
	state.SetTemperature(mc.Temperature)
	state.SetTopP(mc.TopP)
	state.SetTopK(mc.TopK)

	router.GET("/", handlers.RenderIndex)
	router.GET("/state", handlers.GetCurrentState)
	router.GET("/messages", handlers.GetAllMessages)
	router.GET("/message/:id/render", handlers.GetMessage)
	router.GET("/model-platform", handlers.GetModelPlatform)
	router.GET("/model", handlers.GetModelsByPlatform(ctx, mc.Region))
	router.GET("/parameters", handlers.GetParameterState)

	router.POST("/select-model", handlers.SelectModel)
	router.POST("/run", mc.Invoke(ctx))
	router.POST("/parameters", handlers.SetParameterState)
	router.POST("/new-conversation", handlers.CreateConversation)

	router.DELETE("/conversation/:id/delete", handlers.DeleteChat)

	if err := router.Run(":31000"); err != nil {
		log.Fatal(err)
	}
}
