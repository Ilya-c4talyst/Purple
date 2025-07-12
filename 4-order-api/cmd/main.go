package main

import (
	"log"
	"net/http"
	"order-api/config"
	"order-api/db"
	"order-api/internal/products"
	"order-api/migrations"
)

func main() {
	// Конфигурация
	config, err := config.LoadConfig()
	checkErr(err)

	// База данных
	dataBase, err := db.NewDb(config.DbConfig)
	checkErr(err)

	// Миграции
	err = migrations.AutoMigrate(dataBase)
	checkErr(err)

	// Репозитории
	repo := products.NewProductsRepository(dataBase)

	// Роутеры
	router := http.NewServeMux()
	products.NewProductsHandler(router, repo)

	// Запуск сервера
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Server is listening on port 8081")
	err = server.ListenAndServe()
	checkErr(err)
}

// Проверка ошибки
func checkErr(err error) {
	if err != nil {
		log.Fatalf("Shutdown because %e", err)
	}
}
