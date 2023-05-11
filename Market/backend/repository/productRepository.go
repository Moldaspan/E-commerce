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
	var books []models.Product
	r.db.Find(&books)
	return books
}
func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	var book models.Product
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return &models.Product{}, err
	}
	return &book, nil
}
