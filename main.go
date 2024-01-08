package main

import (
	"log"
	"net/http"
)

func init() {
	manager := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}

func main() {
	// Serve on port :8080, fudge yeah hardcoded port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
