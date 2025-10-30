package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize database
	db, err := InitDB("tasks.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Create handler with database
	handler := &TaskHandler{db: db}

	// Set up routes
	http.HandleFunc("/tasks", handler.HandleTasks)
	http.HandleFunc("/tasks/", handler.HandleTaskByID)

	// Start server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
