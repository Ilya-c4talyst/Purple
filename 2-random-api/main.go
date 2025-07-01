package main

import (
	"log"
	"net/http"
	"purple/2-random-api/handlers"
)

func main() {

	// Создание сервера и инит обработчика
	mux := http.NewServeMux()
	handlers.NewRandomHandler(mux)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on port 8080")

	// Запуск сервера
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server crashed with error %e", err)
	}
}
