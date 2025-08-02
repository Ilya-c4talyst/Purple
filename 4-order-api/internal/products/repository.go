package products

import (
	"order-api/db"

	"gorm.io/gorm/clause"
)

type ProductsRepository struct {
	Database *db.Db
}

func NewProductsRepository(dataBase *db.Db) *ProductsRepository {
	return &ProductsRepository{
		Database: dataBase,
	}
}

func (r *ProductsRepository) Create(product *Product) (*Product, error) {
	result := r.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *ProductsRepository) Update(product *Product) (*Product, error) {
	result := r.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *ProductsRepository) Delete(id uint) error {
	result := r.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductsRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := r.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductsRepository) GetAll() (*[]Product, error) {
	var products []Product
	result := r.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}
