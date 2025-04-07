package main

import (
	"database/sql"
	"emoji-webpage/internal"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {

	//Add the database connection logic at the beginning of main() function.
	// This ensures that the database is initialized before handling any HTTP requests.
	err := internal.ConnectDatabase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer internal.DB.Close() /// Ensure the database connection is closed when the program exits

	// Create table if it doesn't exist
	err = internal.InitializeDatabase()
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	err = internal.SeedDatabaseFromJSON("internal/emojis.json") // Load emojis from JSON file
	if err != nil {
		fmt.Println("Failed to seed database:", err)
		return
	}

	// Serve static files (CSS, images) from a "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var emoji string
		var inputText string

		// Template for the error page
		errorTmpl := template.Must(template.ParseFiles("templates/error.html"))
		error500Tmpl := template.Must(template.ParseFiles("templates/error500.html"))

		// Template for the HTML page
		// tmpl:= template.Must(template.ParseFiles("index.html"))
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			error500Tmpl.Execute(w, struct {
				Status  int
				Message string
			}{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
			return
			// fmt.Printf("Error loading main template: %v\n", err)
			// return
		}

		// Handle invalid paths with a custom 404 error page
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			errorTmpl.Execute(w, struct {
				Status  int
				Message string
			}{
				Status:  http.StatusNotFound,
				Message: "Page Not Found",
			})
			return
		}

		// Handle GET requests
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			// Render the template with empty input and output
			tmpl.Execute(w, struct {
				InputText string
				Emoji     string
			}{
				InputText: "",
				Emoji:     "",
			})
			return
		}

		// Handle POST requests
		if r.Method == http.MethodPost {
			r.ParseForm()
			inputText = strings.ToLower(r.PostForm.Get("input-text"))

			// Check if the input text is empty
			if inputText == "" {
				emoji = "⚠️Please enter a valid emoji name"
			} else {
				// Fetch emoji from the database using its name
				symbol, err := internal.GetEmojiByName(inputText)
				if err != nil {
					if err == sql.ErrNoRows {
						// Emoji not found in the database
						emoji = "❓Emoji not found"
					} else {
						// Handle other errors
						fmt.Println("Error fetching emoji:", err)
						w.WriteHeader(http.StatusInternalServerError)
						error500Tmpl.Execute(w, struct {
							Status  int
							Message string
						}{
							Status:  http.StatusInternalServerError,
							Message: "Internal Server Error",
						})
						return
					}
				} else {
					emoji = strings.Join(symbol, "") // Combine all matching emojis into one string
				}
			}
		}
		// Execute the template and pass the emoji to it
		tmpl.Execute(w, struct {
			InputText string
			Emoji     string
		}{
			InputText: inputText,
			Emoji:     emoji})
	})

	// Handle search requests
	http.HandleFunc("/suggest", func(w http.ResponseWriter, r *http.Request) {
		prefix := strings.ToLower(r.URL.Query().Get("prefix"))
		if len(prefix) < 2 {
			fmt.Fprint(w, "[]") // Return empty array for short prefixes
			return
		}

		suggestions, err := internal.GetEmojiSuggestions(prefix[:2])
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Return suggestions as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(suggestions)
	})

	// Handle add new emojis requests
http.HandleFunc("/api/v1/emojis", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse JSON request body
    var newEmoji struct {
        Name   string `json:"name"`
        Symbol string `json:"symbol"`
    }
    err := json.NewDecoder(r.Body).Decode(&newEmoji)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate input
    if newEmoji.Name == "" || newEmoji.Symbol == "" {
        http.Error(w, "Name and symbol are required", http.StatusBadRequest)
        return
    }

    // Insert into database
    err = internal.InsertEmoji(newEmoji.Name, newEmoji.Symbol)
    if err != nil {
        if strings.Contains(err.Error(), "UNIQUE constraint failed") {
            http.Error(w, "Emoji already exists", http.StatusConflict)
        } else {
            http.Error(w, "Failed to add emoji", http.StatusInternalServerError)
        }
        return
    }

    // Return success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "success",
        "message": "Emoji added successfully",
    })
})


	// Start the HTTP server on port 8080
	fmt.Println("Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
