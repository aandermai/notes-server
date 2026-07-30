package main

import (
	"log"
	"net/http"

	"github.com/aandermai/notes-server/internal/handler"
)

func main() {
	http.HandleFunc("/notes/", handler.NotesHandler)

	log.Println("Server is running on port 8080\nOpen in browser: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
