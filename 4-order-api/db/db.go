package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"order-api/config"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *config.DbConfig) (*Db, error) {
	// Регистрация pq.StringArray
	dataBase, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Db{dataBase}, nil
}
