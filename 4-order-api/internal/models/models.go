package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Images      []string `gorm:"type:text[]" json:"images"`
}
