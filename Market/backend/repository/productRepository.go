package repository

import (
	"github.com/Moldaspan/E-commerce/backend/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}
func (r *ProductRepository) GetProducts() []models.Product {
	var products []models.Product
	r.db.Find(&products)
	return products
}
func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	var prod models.Product
	err := r.db.Where("id = ?", id).First(&prod).Error
	if err != nil {
		return &models.Product{}, err
	}
	return &prod, nil
}
