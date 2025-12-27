# Emoji Converter Webpage
Welcome to the Emoji Converter Webpage project! This application allows users to search for emojis by name, add new emojis to the database, and interact with a dynamic webpage. The project is built using Go, SQLite, HTML, CSS, and JavaScript.

## Project Structure
```
Emoji Converter webpage/
├── internal/
│   ├── database.go       # Database connection and operations
│   ├── emojis.json       # Predefined list of emojis in JSON format
│   ├── sample.json       # Predefined list of emojis in JSON format #2
├── static/
│   ├── script.js         # JavaScript for handling search suggestions and form submission
│   ├── backgroundimg.jpg # Background image
│   └── style.css         # CSS for styling the webpage
├── templates/
│   ├── error.html        # Error page template
│   ├── error500.html     # Internal server error template
│   └── index.html        # Main HTML template
├── main.go               # Entry point for the app
├── emoji_converter.db    # SQLite database file
├── Makefile              # Automation commands for building, running, and managing the project
├── go.mod                # Go module file
└── README.md             # Documentation
```

## Features
1. Emoji Search and Conversion:
- Users can search for emojis by typing emoji names into the input field.

- Name suggestions are dynamically displayed as users type.

- Emojis are converted and displayed in real-time.

2. Add New Emojis:
- Users can add new emojis to the database using a form.

- The form includes fields for emoji name and symbol.

- A clear button allows users to reset the form.

3. Dynamic Webpage:
- The webpage is styled with CSS and includes a background image for aesthetics.

- JavaScript handles dynamic interactions, such as clearing fields and updating suggestions.

4. REST API:
- A REST API endpoint (POST /api/v1/emojis) allows adding new emojis programmatically.

5. Error Handling:
- Custom error pages (error.html and error500.html) provide user-friendly feedback.

## Prerequisites
Before running this project, ensure you have the following installed on your system:

1. Go (version 1.20 or higher)

2. SQLite (for database operations)

3. A modern web browser (e.g., Chrome, Firefox)

## Setup Instructions
Step 1: Clone the Repository
```
git clone https://github.com/mareerray/emoji-converter-webpage.git
cd emoji-converter-webpage
```
Step 2: Install Dependencies
```
go mod tidy
```
Step 3: Initialize the Database
Run the following command to create tables and seed data from emojis.json:

```
make init-db
```
Step 4: Start the Server
Start the Go server by running:

```
go run main.go
```
The server will start at http://localhost:8000.

## Usage Instructions
1. Emoji Search:
- Open http://localhost:8000 in your browser.

- Type an emoji name (e.g., "smile") into the input field.

- Suggestions will appear dynamically as you type.

2. Add New Emoji:
- Navigate to the left section of the webpage.

- Enter an emoji name and symbol in the respective fields.

- Click "Add Emoji" to save it to the database.

- Click "Clear Form" to reset the fields.

3. Refresh Input Field:
- Click "Refresh" to clear the input field, suggestions, and output container.

## REST API Documentation
POST /api/v1/emojis
Adds a new emoji to the database.

Request Body:
```
json
{
    "name": "palm_tree",
    "symbol": "🌴"
}
```
Response:
201 Created: Emoji added successfully.

```
json
{
    "status": "success",
    "message": "Emoji added successfully"
}
```
400 Bad Request: Invalid request body.

409 Conflict: Emoji already exists.

## Development Workflow
Using Makefile Commands:
The Makefile provides automation commands for common tasks:

Command	        Description
make init-db	Initializes the database tables
make seed-db	Seeds data into the database
make run	    Runs the Go server
make clean	    Removes build artifacts and database

## Tables in the Database
1. emojis:

- id (Primary Key)

- name (Unique)

- symbol

2. emoji_prefixes:

- prefix

- emoji_name (Foreign Key referencing emojis.name)

### Relationships
emojis → emoji_prefixes: One-to-Many (1:N) where each emoji can have multiple prefixes.

## Technologies Used
1. Backend:

- Go (Golang) for server-side logic.

- SQLite for database operations.

2. Frontend:

- HTML for structure.

- CSS for styling.

- JavaScript for dynamic interactions.

3. Other Tools:

- Makefile for automation.

- JSON files (emojis.json, sample.json) for initial data.

## Future Enhancements
- Add user authentication to track personalized emoji usage.

- Implement additional REST API endpoints (e.g., GET /api/v1/emojis/{name}).

- Improve responsiveness with advanced CSS techniques (e.g., grid layout).

- Add analytics to track popular emoji searches.


## Contributor
For questions or feedback, feel free to reach out:

Mayuree Reunsati

GitHub: https://github.com/mareerray
