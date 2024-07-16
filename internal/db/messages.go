package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InsertMessage inserts a new message into the database
func InsertMessage(conversationID int, human, response, platform, model string) error {
	stmt, err := DB.Prepare("INSERT INTO messages(conversation_id, human, response, platform, model) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(conversationID, human, response, platform, model)
	return err
}

// GetMessagesByConversationID returns all messages in the specified conversation
func GetMessagesByConversationID(conversationID int) ([]Message, error) {
	rows, err := DB.Query(`
		SELECT id, conversation_id, COALESCE(human, ''), COALESCE(response, ''), platform, model 
		FROM messages 
		WHERE conversation_id = ? 
		ORDER BY id ASC`, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.ConversationID, &msg.Human, &msg.Response, &msg.Platform, &msg.Model); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

// GetMessageByID returns a message with the specified ID
func GetMessageByID(id int) (Message, error) {
	var msg Message
	err := DB.QueryRow("SELECT id, conversation_id, human, response, platform, model FROM messages WHERE id = ?", id).Scan(&msg.ID, &msg.ConversationID, &msg.Human, &msg.Response, &msg.Platform, &msg.Model)
	return msg, err
}
