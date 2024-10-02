package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupApi()

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}

func setupApi() {
	ctx := context.Background()

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login", manager.loginHandler)
}
