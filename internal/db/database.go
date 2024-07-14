package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	Conn *sql.DB
}

// NewDB creates a new database if it doesn't exist
// and returns a connection to it
func NewDB() (Client, error) {
	var output Client
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./chat.db")
	if err != nil {
		return output, err
	}
	output.Conn = db

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
		return output, err
	}
	return output, nil
}
