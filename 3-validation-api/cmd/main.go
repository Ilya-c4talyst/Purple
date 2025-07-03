package main

import (
	"log"
	"net/http"
	"purple/validation/config"
	"purple/validation/internal/verify"
)

func main() {

	config := config.NewConfig()

	router := http.NewServeMux()
	verify.NewEmailHandler(router, config)

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
