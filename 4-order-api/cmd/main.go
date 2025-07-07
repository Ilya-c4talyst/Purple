package main

import (
	"log"
	"order-api/config"
	"order-api/db"
	"order-api/migrations"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Shutdown because %e", err)
	}

	dataBase, err := db.NewDb(config.DbConfig)

	if err != nil {
		log.Fatalf("Shutdown because %e", err)
	}

	migrations.AutoMigrate(dataBase)

}
