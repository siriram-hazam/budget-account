package main

import (
	"log"
	"net/http"

	"github.com/siriramhazam/budget-authen/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	authHandler := handler.NewAuthHandler()
	mux.HandleFunc("/auth", authHandler.Login)

	log.Println("Server stating at :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
