package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbConfig *DbConfig
}

type DbConfig struct {
	Host     string
	Username string
	Password string
	DbName   string
	Port     string
	SslMode  string
	DSN      string
}

// Загружаем переменные из .env
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки .env файла: %v", err)
	}

	host := os.Getenv("host")
	username := os.Getenv("user")
	password := os.Getenv("password")
	dbName := os.Getenv("dbname")
	port := os.Getenv("port")
	sslMode := os.Getenv("sslmode")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, username, password, dbName, port, sslMode,
	)

	cfg := &Config{
		DbConfig: &DbConfig{
			Host:     os.Getenv("host"),
			Username: os.Getenv("user"),
			Password: os.Getenv("password"),
			DbName:   os.Getenv("dbname"),
			Port:     os.Getenv("port"),
			SslMode:  os.Getenv("sslmode"),
			DSN:      dsn,
		},
	}

	// Простая проверка, что все значения загружены
	if cfg.DbConfig.Host == "" || cfg.DbConfig.Username == "" || cfg.DbConfig.Password == "" ||
		cfg.DbConfig.DbName == "" || cfg.DbConfig.Port == "" {
		return nil, fmt.Errorf("не все переменные окружения установлены")
	}

	return cfg, nil
}
