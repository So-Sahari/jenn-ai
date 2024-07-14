package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func (c *Client) InsertMessage(conversationID int, human, response, platform, model string) error {
	stmt, err := c.Conn.Prepare("INSERT INTO messages(conversation_id, human, response, platform, model) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(conversationID, human, response, platform, model)
	return err
}

func (c *Client) GetMessagesByConversationID(conversationID int) ([]Message, error) {
	rows, err := c.Conn.Query(`
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

func (c *Client) GetMessageByID(id int) (Message, error) {
	var msg Message
	err := c.Conn.QueryRow("SELECT id, conversation_id, human, response, platform, model FROM messages WHERE id = ?", id).Scan(&msg.ID, &msg.ConversationID, &msg.Human, &msg.Response, &msg.Platform, &msg.Model)
	return msg, err
}
