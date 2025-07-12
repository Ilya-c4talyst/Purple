package products

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`
}

type ProductCreateRequest struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
}

func NewProduct(name string, description string, images pq.StringArray) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
