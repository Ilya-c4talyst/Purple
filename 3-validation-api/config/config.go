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
		log.Println("Error while getting .env")
	}

	config := &Config{
		Email: &EmailConfig{
			os.Getenv("EMAIL"),
			os.Getenv("PASSWORD"),
			os.Getenv("ADDRESS"),
		},
	}

	return config
}
