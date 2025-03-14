package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type Message struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Text string `json:"text"`
}

// Initialize in-memory storage and service as global variables

var (
	messages []Message
)

func init() {
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
}

// Handlers/controllers

func getMessages(c echo.Context) error {
	log.Println("getMessages called")
	
	time.Sleep(2 * time.Second);  // 
	return c.JSON(http.StatusOK, messages)
}

// Main function with routes

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/messages", getMessages)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:8081"))
}