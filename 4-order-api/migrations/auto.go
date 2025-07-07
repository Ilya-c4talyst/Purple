package migrations

import (
	"order-api/db"
	"order-api/internal/models"
)

func AutoMigrate(db *db.Db) error {
	err := db.AutoMigrate(&models.Product{})
	return err
}
