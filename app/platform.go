package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var modelSource = []string{
	"Bedrock",
	"Ollama",
}

func getModelPlatform(c *gin.Context) {
	modelPlatform := "<option disabled selected>Model Platform</option>"
	for _, plat := range modelSource {
		modelPlatform += fmt.Sprintf("<option>%s</option>", plat)
	}
	c.Data(http.StatusOK, "text/html", []byte(modelPlatform))
	selectedOption := c.PostForm("platform-option")
	fmt.Printf("Received the following: %s\n", selectedOption)
}
