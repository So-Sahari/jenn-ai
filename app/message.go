package app

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"jenn-ai/internal/parser"
	"jenn-ai/internal/state"

	"github.com/gin-gonic/gin"
)

type ChatMessage struct {
	Human    template.HTML
	Response template.HTML
	Platform string
	Model    string
}

func createConversation(c *gin.Context) {
	conversationID, err := createNewConversation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new conversation"})
		return
	}

	appState := state.GetState()
	appState.SetConversationID(conversationID)

	c.HTML(http.StatusOK, "chat.html", gin.H{
		"ChatMessages": []ChatMessage{},
		"Platform":     appState.GetPlatform(),
		"Model":        appState.GetModel(),
	})
}

func getMessagesFromDB(c *gin.Context) {
	conversationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
		return
	}

	messages, err := getMessagesByConversationID(conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	appState := state.GetState()
	if len(messages) > 0 {
		lastMessage := messages[len(messages)-1]
		appState.SetPlatform(lastMessage.Platform)
		appState.SetModel(lastMessage.Model)
		appState.SetConversationID(lastMessage.ConversationID)
	}

	var chatMessages []ChatMessage
	for _, msg := range messages {
		// parse markdown
		parsed, err := parser.ParseMD(msg.Response)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		parsed = strings.ReplaceAll(parsed, "<pre>", "<div class='card bg-base-100 shadow-xl'><div class='card-body overflow-x-auto'><pre>")
		parsed = strings.ReplaceAll(parsed, "</pre>", "</pre></div></div>")

		chatMessages = append(chatMessages, ChatMessage{
			Human:    template.HTML(msg.Human),
			Response: template.HTML(parsed),
			Platform: msg.Platform,
			Model:    msg.Model,
		})
	}

	c.HTML(http.StatusOK, "chat.html", gin.H{
		"ChatMessages": chatMessages,
	})
}

func getAllMessagesFromDB(c *gin.Context) {
	conversations, err := getAllConversations()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "sidebar.html", gin.H{
		"Conversations": conversations,
	})
}

func getAllConversationsHandler(c *gin.Context) {
	conversations, err := getAllConversations()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "sidebar.html", gin.H{
		"Conversations": conversations,
	})
}
