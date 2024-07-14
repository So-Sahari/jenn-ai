package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"jenn-ai/internal/config"
)

func GetModelPlatform(c *gin.Context) {
	modelPlatform := "<option disabled selected>Model Platform</option>"
	for _, plat := range config.SupportedPlatforms {
		modelPlatform += fmt.Sprintf("<option>%s</option>", plat)
	}
	c.Data(http.StatusOK, "text/html", []byte(modelPlatform))
	selectedOption := c.PostForm("platform-option")
	fmt.Printf("Received the following: %s\n", selectedOption)
}
