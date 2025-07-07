package db

import (
	"order-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *config.DbConfig) (*Db, error) {
	dataBase, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Db{dataBase}, nil
}
