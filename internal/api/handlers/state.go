package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"jenn-ai/internal/state"
)

func GetCurrentState(c *gin.Context) {
	appState := state.GetState()
	model := appState.GetModel()
	if model == "" {
		model = "No Model Selected"
	}

	c.HTML(http.StatusOK, "state.html", gin.H{
		"Platform": appState.GetPlatform(),
		"Model":    model,
	})
}

func GetParameterState(c *gin.Context) {
	appState := state.GetState()

	c.HTML(http.StatusOK, "parameters.html", gin.H{
		"Tokens":      appState.GetMaxTokens(),
		"TopP":        appState.GetTopP(),
		"TopK":        appState.GetTopK(),
		"Temperature": appState.GetTemperature(),
	})
}

func SetParameterState(c *gin.Context) {
	appState := state.GetState()
	maxTokens := c.PostForm("maxTokens")
	if maxTokens != "" {
		maxTokensInt, err := strconv.Atoi(maxTokens)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set max tokens"})
			return
		}
		appState.SetMaxTokens(maxTokensInt)
	}

	topK := c.PostForm("topK")
	if topK != "" {
		topKInt, err := strconv.Atoi(topK)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set top k"})
			return
		}
		appState.SetTopK(topKInt)
	}

	topP := c.PostForm("topP")
	if topP != "" {
		topPFloat, err := strconv.ParseFloat(topP, 64)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set top p"})
			return
		}
		appState.SetTopP(topPFloat)
	}

	temperature := c.PostForm("temperature")
	if temperature != "" {
		temperatureFloat, err := strconv.ParseFloat(temperature, 64)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set temperature"})
			return
		}
		appState.SetTemperature(temperatureFloat)
	}

	c.HTML(http.StatusOK, "parameters.html", gin.H{
		"Tokens":      appState.GetMaxTokens(),
		"TopP":        appState.GetTopP(),
		"TopK":        appState.GetTopK(),
		"Temperature": appState.GetTemperature(),
	})
}
