package app

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS conversations (
		id INTEGER PRIMARY KEY AUTOINCREMENT
	);
	
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		conversation_id INTEGER NOT NULL,
		human TEXT,
		response TEXT,
		platform TEXT,
		model TEXT,
		FOREIGN KEY (conversation_id) REFERENCES conversations(id)
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func insertMessage(conversationID int, human, response, platform, model string) error {
	stmt, err := db.Prepare("INSERT INTO messages(conversation_id, human, response, platform, model) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(conversationID, human, response, platform, model)
	return err
}

func getMessagesByConversationID(conversationID int) ([]Message, error) {
	rows, err := db.Query("SELECT id, conversation_id, human, response, platform, model FROM messages WHERE conversation_id = ?", conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.ConversationID, &msg.Human, &msg.Response, &msg.Platform, &msg.Model); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func getAllConversations() ([]Conversation, error) {
	query := `
		SELECT c.id, m.human, m.response
		FROM conversations c
		LEFT JOIN messages m ON m.id = (
			SELECT id FROM messages
			WHERE conversation_id = c.id
			ORDER BY id DESC
			LIMIT 1
		)
		ORDER BY c.id
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query %s: %v", query, err)
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
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}
	return conversations, nil
}

func getMessageByID(id int) (Message, error) {
	var msg Message
	err := db.QueryRow("SELECT id, conversation_id, human, response, platform, model FROM messages WHERE id = ?", id).Scan(&msg.ID, &msg.ConversationID, &msg.Human, &msg.Response, &msg.Platform, &msg.Model)
	return msg, err
}

func createNewConversation() (int, error) {
	stmt, err := db.Prepare("INSERT INTO conversations DEFAULT VALUES RETURNING id")
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
