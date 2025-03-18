package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Message struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}

type CircuitBreaker struct {
	state       int32
	failures    int32
	lock        sync.Mutex
	lastFailure time.Time
}

const (
	Closed     = 0
	Open       = 1
	HalfOpen   = 2
	MaxRetries = 3
	OpenTime   = 10 * time.Second
)

var cb CircuitBreaker

func fetchMessages() ([]Message, error) {
	url := "http://localhost:8081/messages/flaky"
	var messages []Message

	for i := 0; i < MaxRetries; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			json.NewDecoder(resp.Body).Decode(&messages)
			log.Println("Messages successfully fetched!")
			return messages, nil
		}

		cb.lock.Lock()
		atomic.AddInt32(&cb.failures, 1) // Count request failures (too many opens the CB)
		cb.lock.Unlock()

		log.Println("Fetch failed, retrying in 1s...")
		time.Sleep(1 * time.Second)
	}

	return nil, http.ErrServerClosed
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	cb.lock.Lock()
	if atomic.LoadInt32(&cb.state) == Open {
		if time.Since(cb.lastFailure) > OpenTime {
			log.Println("[Circuit Breaker] Moving to HALF-OPEN state.")
			atomic.StoreInt32(&cb.state, HalfOpen)
		}
		cb.lock.Unlock()
	} else {
		cb.lock.Unlock()
	}

	if atomic.LoadInt32(&cb.state) == Open {
		log.Println("[Circuit Breaker] OPEN - Returning fallback response.")
		sendJSON(w, []Message{{ID: 0, User: "System", Text: "Service unavailable. Please try again later. (FALLBACK - CIRCUIT OPEN)"}})
		return
	}

	messages, err := fetchMessages()
	if err != nil {
		cb.lock.Lock()
		if cb.failures >= MaxRetries {
			log.Println("[Circuit Breaker] Too many failures, switching to OPEN state.")
			atomic.StoreInt32(&cb.state, Open)
			cb.lastFailure = time.Now()
		}
		cb.lock.Unlock()

		sendJSON(w, []Message{{ID: 0, User: "System", Text: "Temporary failure. Please try again later. (BACKEND ERROR)"}})
		return
	}

	// Success after being open: we can close back the circuit. (=> reset nb of failures)
	cb.lock.Lock()
	if atomic.LoadInt32(&cb.state) == HalfOpen {
		log.Println("[Circuit Breaker] HALF-OPEN -> CLOSED - Service recovered.")
		atomic.StoreInt32(&cb.state, Closed)
	}
	cb.failures = 0
	cb.lock.Unlock()

	sendJSON(w, messages)
}

func sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", getMessages)
	log.Println("Fetcher running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
