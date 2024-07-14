package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func (c *Client) GetAllConversations() ([]Conversation, error) {
	rows, err := c.Conn.Query(`
		SELECT c.id, COALESCE(m.human, ''), COALESCE(m.response, '')
		FROM conversations c
		LEFT JOIN messages m ON m.id = (
			SELECT id FROM messages
			WHERE conversation_id = c.id
			ORDER BY id DESC
			LIMIT 1
		)
		ORDER BY c.id
	`)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var conversations []Conversation
	for rows.Next() {
		var conv Conversation
		if err := rows.Scan(&conv.ID, &conv.LatestHuman, &conv.LatestResponse); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		conversations = append(conversations, conv)
	}
	return conversations, nil
}

func (c *Client) CreateNewConversation() (int, error) {
	stmt, err := c.Conn.Prepare("INSERT INTO conversations DEFAULT VALUES RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var conversationID int
	err = stmt.QueryRow().Scan(&conversationID)
	if err != nil {
		return 0, err
	}
	return conversationID, nil
}

func (c *Client) DeleteConversation(conversationID int) error {
	// Delete associated messages first
	_, err := c.Conn.Exec("DELETE FROM messages WHERE conversation_id = ?", conversationID)
	if err != nil {
		return err
	}

	// Delete the conversation
	_, err = c.Conn.Exec("DELETE FROM conversations WHERE id = ?", conversationID)
	return err
}
