package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// NewDB creates a new database if it doesn't exist
// and returns an error if it fails
func NewDB(dbPath string) error {
	var err error

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
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
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

// CloseDB closes the database
func CloseDB() {
	DB.Close()
}
