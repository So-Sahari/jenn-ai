package db

type Message struct {
	ID             int    `json:"id"`
	ConversationID int    `json:"conversation_id"`
	Human          string `json:"human"`
	Response       string `json:"response"`
	Platform       string `json:"platform"`
	Model          string `json:"model"`
}

type Conversation struct {
	ID             int    `json:"id"`
	LatestHuman    string `json:"latest_human"`
	LatestResponse string `json:"latest_response"`
}
