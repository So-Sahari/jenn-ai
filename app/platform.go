package app

import (
	"fmt"
	"net/http"

	"jenn-ai/internal/fuzzy"

	"github.com/gin-gonic/gin"
)

var platform string

func getModelPlatform(c *gin.Context) {
	var modelPlatform string

	for _, plat := range fuzzy.ModelSource {
		modelPlatform += fmt.Sprintf(`<a href="#" 
      data-value="%s"
      hx-post="/select-platform?option=%s" 
      class="platform-item block px-4 py-2 text-sm text-black-700 hover:bg-blue-400">
      %s</a>`, plat, plat, plat)
	}
	c.Data(http.StatusOK, "text/html", []byte(modelPlatform))
}

func setModelPlatform(c *gin.Context) {
	selectedOption := c.Query("option")
	fmt.Printf("Received the following: %s\n", selectedOption)
	platform = selectedOption
}
