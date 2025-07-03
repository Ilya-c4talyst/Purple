package main

import (
	"log"
	"net/http"
	"purple/email_api/config"
	"purple/email_api/internal/email"
)

func main() {

	config := config.NewConfig()

	router := http.NewServeMux()
	email.NewEmailHandler(router, config)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Error %e", err)
	}
}
