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

	// Try fetching messages
	messages, err := fetchMessages()
	data := struct {
		Messages []Message
		Error    bool
	}{
		Messages: messages,
		Error:    err != nil,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Frontend running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
