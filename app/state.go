package app

import (
	"net/http"

	"jenn-ai/internal/state"

	"github.com/gin-gonic/gin"
)

func getCurrentState(c *gin.Context) {
	appState := state.GetState()
	c.JSON(http.StatusOK, gin.H{
		"platform": appState.GetPlatform(),
		"model":    appState.GetModel(),
	})
}
