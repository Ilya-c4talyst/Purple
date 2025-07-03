package main

import (
	"log"
	"net/http"
	"purple/validation/config"
	"purple/validation/internal/email"
)

func main() {

	config := config.NewConfig()

	router := http.NewServeMux()
	email.NewEmailHandler(router, config)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Error %e", err)
	}
}
