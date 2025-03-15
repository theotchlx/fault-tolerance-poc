package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

type Message struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}

// Declare in-memory storage as global variable

var (
	messages []Message
)

// Service functions

func getMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("-- getMessages called --")

	time.Sleep(2 * time.Second) // Simulate processing delay

	sendJSON(w, messages)
}

// Simulate failure: slow response / delay (5 seconds)
func slowResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("slowResponse called")

	time.Sleep(18 * time.Second) // Total 20 seconds.
	getMessages(w, r)
}

// To test retry, timeout, fallback...

// Always fails
func alwaysFail(w http.ResponseWriter, r *http.Request) {
	log.Println("alwaysFail called")

	http.Error(w, "Service unavailable", http.StatusInternalServerError)
}

// To test fallback, circuit breaker...

// Random failures (50% chance)
func unreliableResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("unreliableResponse called")

	if rand.Float32() < 0.5 {
		http.Error(w, "Random failure", http.StatusInternalServerError)
		return
	}
	getMessages(w, r)
}

// To test retry, circuit breaker...

// First 2 requests fail, then succeed
var flakyCounter int32

func flakyResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("flakyResponse called")

	if atomic.AddInt32(&flakyCounter, 1) <= 2 {
		http.Error(w, "Temporary failure", http.StatusInternalServerError)
		return
	}
	getMessages(w, r)
}

// Also to test retry, circuit breaker...

// Returns a default fallback message when failing
func fallbackResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("fallbackResponse called")

	time.Sleep(1 * time.Second) // Simulate processing delay
	sendJSON(w, []Message{
		{ID: 0, User: "System", Text: "This is a fallback message that is detected and handled to create a seamless user experience."},
	})
}

// To test fallback

// Helper function to send JSON responses
func sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Main function with routes

func main() {

	messages = []Message{
		{ID: 1, User: "Teal", Text: "There is something inside me, and they don't know if there is a cure."},
		{ID: 2, User: "Perpetua", Text: "A demonic possession, unlike any before."},
		{ID: 3, User: "Teal", Text: "It's a sickening heartache, and it's slowly tormenting my soul."},
		{ID: 4, User: "Perpetua", Text: "I've invested my prayers, into making me whole."},
		{ID: 5, User: "Teal", Text: "It's a sickening heartache, and it's slowly tormenting my soul,"},
		{ID: 6, User: "Teal", Text: "I should have known, not to give in, I should have known, not to give in."},
		{ID: 7, User: "Perpetua", Text: "Blasphemy, heresy, save me from the monster that is eating me,"},
		{ID: 8, User: "Perpetua", Text: "I'm victimized, blasphemy, heresy,"},
		{ID: 9, User: "Perpetua", Text: "Save me, from the bottom of my heart I know, I'm satanized, I'm satanized, I'm satanized."},
		{ID: 10, User: "Teal", Text: "An nescitis quoniam membra vestra"},
		{ID: 11, User: "Teal", Text: "Templum est Spiritus Sancti"},
		{ID: 12, User: "Teal", Text: "Qui in vobis est"},
		{ID: 13, User: "Teal", Text: "Quem habetis a Deo?"},
		{ID: 14, User: "Teal", Text: "Et non estis vestri"},
		{ID: 15, User: "Perpetua", Text: "From the bottom of my heart I know, I'm satanized."},
	}

	rand.Seed(time.Now().UnixNano())

	// Routes
	http.HandleFunc("/messages", getMessages) // Normal messages endpoint. No expected failures.
	http.HandleFunc("/messages/slow", slowResponse)
	http.HandleFunc("/messages/down", alwaysFail)
	http.HandleFunc("/messages/unreliable", unreliableResponse)
	http.HandleFunc("/messages/flaky", flakyResponse)
	http.HandleFunc("/messages/fallback", fallbackResponse)

	// Start server
	log.Println("Backend running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
