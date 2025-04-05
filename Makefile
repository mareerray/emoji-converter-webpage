# Define variables
APP_NAME = emoji-webpage
DB_FILE = emoji_converter.db

# Default target (runs the application)
run:
	go run main.go

# Build the application binary
build:
	go build -o $(APP_NAME)

# Test the application
test:
	go test ./...

# Clean up build artifacts and database file
clean:
	rm -f $(APP_NAME) $(DB_FILE)

# Initialize the database (create if it doesn't exist)
init-db:
	sqlite3 $(DB_FILE) "\
	CREATE TABLE IF NOT EXISTS emojis (\
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL, 
		symbol TEXT NOT NULL \
	);\
	CREATE TABLE IF NOT EXISTS emoji_prefixes ( \
			prefix TEXT NOT NULL, \
			emoji_name TEXT NOT NULL, \
			FOREIGN KEY (emoji_name) REFERENCES emojis(name) ON DELETE CASCADE \
	); \
	"
seed-db:
	sqlite3 $(DB_FILE) " \
		INSERT INTO emojis (name, symbol) VALUES \
		('smile', '😄'), \
		('heart', '❤️'), \
		... ALL YOUR EMOJIS ... ; \
		INSERT INTO emoji_prefixes (prefix, emoji_name) VALUES \
		('sm', 'smile'), \
		('he', 'heart'), \
		... ALL PREFIXES ... ; \
	"

# Format Go code
fmt:
	go fmt ./...

# Tidy up dependencies
tidy:
	go mod tidy

# Help documentation for Makefile commands
help:
	@echo "Available commands:"
	@echo "  make run       - Run the application"
	@echo "  make build     - Build the application binary"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Remove build artifacts and database file"
	@echo "  make init-db   - Initialize the SQLite database"
	@echo "  make fmt       - Format Go code"
	@echo "  make tidy      - Tidy up dependencies"
