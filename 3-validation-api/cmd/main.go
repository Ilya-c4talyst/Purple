package main

import (
	"3-validation-api/config"
	"3-validation-api/internal/verify"
	"log"
	"net/http"
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
