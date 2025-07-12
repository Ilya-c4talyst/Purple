package migrations

import (
	"order-api/db"
	"order-api/internal/products"
)

func AutoMigrate(db *db.Db) error {
	err := db.AutoMigrate(&products.Product{})
	return err
}
