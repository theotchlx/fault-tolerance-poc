package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// Message struct
type Message struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
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
		args := os.Args
		var endpoint string
		if len(args) != 2 {
			endpoint = ""
		} else {
			endpoint = args[1]
		}

		resp, err := http.Get("http://localhost:8081/messages" + endpoint)

		if err != nil {
			http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, resp.Body) // Stream response, no buffer
	})

	log.Println("Frontend running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
