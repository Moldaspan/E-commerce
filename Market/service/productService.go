package service

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/repositories"
)

type ProductServiceInterface interface {
	CreateProduct(*models.Product) error
	GetProductByID(uint) (*models.Product, error)
	UpdateProduct(*models.Product) error
	DeleteProduct(uint) error
	GetProducts() ([]models.Product, error)
	GetCommentsByProductId(uint) ([]*models.Comment, error)
	GetProductAverageRating(uint) (float32, error)
	SearchByName(name string) ([]models.Product, error)
	SearchByPriceRange(minPrice, maxPrice float64) ([]models.Product, error)
}

type CategoryServiceInterface interface {
	CreateCategory(category *models.Category) error
	GetCategoryByID(uint) (*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(uint) error
	GetCategories() ([]models.Category, error)
}

type ProductServiceV1 struct {
	productRepos repositories.ProductRepositoryInterface
}

type CategoryServiceV1 struct {
	categoryRepos repositories.CategoryRepositoryInterface
}

func NewProductService() *ProductServiceV1 {
	return &ProductServiceV1{productRepos: repositories.NewProductRepository()}
}

func NewCategoryService() *CategoryServiceV1 {
	return &CategoryServiceV1{categoryRepos: repositories.NewCategoryRepository()}
}

func (p ProductServiceV1) GetProductAverageRating(id uint) (float32, error) {
	return p.productRepos.GetProductAverageRating(id)
}

func (p ProductServiceV1) GetProducts() ([]models.Product, error) {
	return p.productRepos.GetProducts()
}

func (p ProductServiceV1) SearchByName(title string) ([]models.Product, error) {
	return p.productRepos.SearchByName(title)
}

func (p ProductServiceV1) SearchByPriceRange(minPrice, maxPrice float64) ([]models.Product, error) {
	return p.productRepos.SearchByPriceRange(minPrice, maxPrice)
}

func (p ProductServiceV1) GetCommentsByProductId(id uint) ([]*models.Comment, error) {
	return p.productRepos.GetCommentsByProductId(id)
}

func (p ProductServiceV1) CreateProduct(product *models.Product) error {
	return p.productRepos.CreateProduct(product)
}

func (p ProductServiceV1) GetProductByID(id uint) (*models.Product, error) {
	return p.productRepos.GetProductByID(id)
}

func (p ProductServiceV1) UpdateProduct(product *models.Product) error {
	return p.productRepos.UpdateProduct(product)
}

func (p ProductServiceV1) DeleteProduct(id uint) error {
	return p.productRepos.DeleteProduct(id)
}

func (c CategoryServiceV1) CreateCategory(category *models.Category) error {
	return c.categoryRepos.CreateCategory(category)
}

func (c CategoryServiceV1) GetCategoryByID(id uint) (*models.Category, error) {
	return c.categoryRepos.GetCategoryByID(id)
}

func (c CategoryServiceV1) UpdateCategory(category *models.Category) error {
	return c.categoryRepos.UpdateCategory(category)
}

func (c CategoryServiceV1) DeleteCategory(id uint) error {
	return c.categoryRepos.DeleteCategory(id)
}

func (c CategoryServiceV1) GetCategories() ([]models.Category, error) {
	return c.categoryRepos.GetCategories()
}
