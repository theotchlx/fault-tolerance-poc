package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Message struct
type Message struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}

// Fetch messages from backend
func fetchMessages() ([]Message, error) {
	resp, err := http.Get("http://localhost:8081/messages")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var messages []Message
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// Serve frontend
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")

	tmpl.Execute(w, nil)
}

func main() {
	// Register the file server handler for /pictures/
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures/"))))
	// Note that this makes the pictures folder accessible at http://localhost:8080/pictures/

	// Register the homePage handler for the root path
	http.HandleFunc("/", homePage)

	// Fetch messages endpoint for the JS to fetch from
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/messages")
		if err != nil {
			http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, resp.Body) // Stream response
	})

	log.Println("Frontend running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
