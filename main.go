package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Get MongoDB URI from environment or use default
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Initialize database
	db, err := InitDB(mongoURI)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	log.Println("Connected to MongoDB")

	// Create handler with database
	handler := &TaskHandler{db: db}

	// Set up routes
	http.HandleFunc("/tasks", handler.HandleTasks)
	http.HandleFunc("/tasks/", handler.HandleTaskByID)

	// Start server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
