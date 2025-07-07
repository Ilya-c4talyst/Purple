package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email *EmailConfig
}

type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	address := os.Getenv("ADDRESS")

	if email == "" {
		email = "default@example.com"
	}
	if password == "" {
		password = "default_password"
	}
	if address == "" {
		address = "smtp.yandex.ru"
	}

	config := &Config{
		Email: &EmailConfig{
			Email:    email,
			Password: password,
			Address:  address,
		},
	}

	return config
}
