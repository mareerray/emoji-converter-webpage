package internal

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// ConnectDatabase establishes a connection to the SQLite database
func ConnectDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "file:emoji_converter.db?_utf8=1")
	if err != nil {
		return err
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection established")
	return nil
}

// Emoji represents the structure of an emoji in the JSON file
type Emoji struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// InitializeDatabase sets up the database using the schema file
func InitializeDatabase() error {
    schemaSQL, err := os.ReadFile("schema.sql")
    if err != nil {
        return err
    }

    _, err = DB.Exec(string(schemaSQL))
    if err != nil {
        return err
    }

    log.Println("Database schema initialized")
    return nil
}


// SeedDatabaseFromJSON loads emojis from a JSON file and inserts them into the database
func SeedDatabaseFromJSON(filePath string) error {
	// Open the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Parse the JSON data into a slice of Emoji structs
	var emojis []Emoji
	err = json.Unmarshal(data, &emojis)
	if err != nil {
		return err
	}

	// Insert each emoji into the database
	for _, emoji := range emojis {
		// fmt.Printf("Inserting: %s -> %s\n", emoji.Name, emoji.Symbol) // Debug log
		// Insert into emojis table
		_, err := DB.Exec("INSERT INTO emojis (name, symbol) VALUES (?, ?)", emoji.Name, emoji.Symbol)
		if err != nil {
			return err
		}

		// Insert prefix into emoji_prefixes table
		prefix := emoji.Name[:2] // First 2 characters of the emoji name
		_, err = DB.Exec("INSERT INTO emoji_prefixes (prefix, emoji_name) VALUES (?, ?)", prefix, emoji.Name)
		if err != nil {
			return err
		}
	}
	log.Println("Database seeded with emojis from JSON file")
	return nil
}

// CreateTable creates the emoji table if it doesn't exist
// func CreateTable() error {
// 	// Drop existing tables (useful during development)
// 	_, _ = DB.Exec("DROP TABLE IF EXISTS emoji_prefixes;")
// 	_, _ = DB.Exec("DROP TABLE IF EXISTS emojis;")

// 	// Create emojis table with UTF-8 collation
// 	_, err := DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS emojis (
// 			id INTEGER PRIMARY KEY AUTOINCREMENT, 
// 			name TEXT NOT NULL UNIQUE COLLATE NOCASE, 
// 			symbol TEXT NOT NULL COLLATE NOCASE
// 		);
// 	`)
// 	if err != nil {
// 		return err
// 	}

// 	// Create emoji_prefixes table
// 	_, err = DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS emoji_prefixes (
// 			prefix TEXT NOT NULL, 
// 			emoji_name TEXT NOT NULL, 
// 			FOREIGN KEY (emoji_name) REFERENCES emojis(name) ON DELETE CASCADE
// 			);
// 		`)
// 	return err
// }

func InsertEmoji(name string, symbol string) error {
	// Insert into emojis table
	_, err := DB.Exec("INSERT INTO emojis (name, symbol) VALUES (?, ?)", name, symbol)
	if err != nil {
		return err
	}

	// Generate prefixes (e.g., first 2 characters)
	prefix := name[:2]
	_, err = DB.Exec("INSERT INTO emoji_prefixes (prefix, emoji_name) VALUES (?, ?)", prefix, name)
	return err
}

func GetEmojiByName(name string) ([]string, error) {
	query := `
	SELECT symbol FROM emojis
	WHERE name = ?`

	rows, err := DB.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var symbols []string
	for rows.Next() {
		var symbol string
		err := rows.Scan(&symbol)
		if err != nil {
			return nil, err
		}
		symbols = append(symbols, symbol)
	}

	if len(symbols) == 0 {
		return nil, sql.ErrNoRows // Return an error if no results are found
	}

	return symbols, nil
}

// GetEmojiSuggestions returns emoji names matching the prefix
func GetEmojiSuggestions(prefix string) ([]string, error) {
	// rows, err := DB.Query("SELECT emoji_name FROM emoji_prefixes WHERE prefix = ?", prefix) //prevent display of the same emoji
	rows, err := DB.Query("SELECT name FROM emojis WHERE name LIKE ? || '%'", prefix)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suggestions []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		suggestions = append(suggestions, name)
	}

	return suggestions, nil
}
