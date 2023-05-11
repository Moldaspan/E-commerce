package repositories

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/settings"
	"gorm.io/gorm"
	"log"
)

type ProductRepositoryInterface interface {
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
type CategoryRepositoryInterface interface {
	CreateCategory(category *models.Category) error
	GetCategoryByID(uint) (*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(uint) error
	GetCategories() ([]models.Category, error)
}

type ProductRepositoryV1 struct {
	DB *gorm.DB
}

type CategoryRepositoryV1 struct {
	DB *gorm.DB
}

func NewProductRepository() *ProductRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &ProductRepositoryV1{}
	}
	return &ProductRepositoryV1{DB: db}
}

func (p *ProductRepositoryV1) GetProductAverageRating(id uint) (float32, error) {
	var ratingAvg float32
	result := p.DB.Table("ratings").
		Select("ROUND(AVG(value)) as rating_average").
		Joins("JOIN products on ratings.product_id = products.id").
		Where("products.id = ?", id).
		Scan(&ratingAvg)
	if result.Error != nil {
		return -1, nil
	}
	return ratingAvg, nil
}

func (p *ProductRepositoryV1) CreateProduct(product *models.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepositoryV1) UpdateProduct(product *models.Product) error {
	return p.DB.Save(product).Error
}

func (p *ProductRepositoryV1) DeleteProduct(id uint) error {
	return p.DB.Delete(&models.Product{}, id).Error
}

func (p *ProductRepositoryV1) GetProducts() ([]models.Product, error) {
	products := make([]models.Product, 0)

	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (ps ProductRepositoryV1) SearchByName(title string) ([]models.Product, error) {
	var products []models.Product
	err := ps.DB.Where("title LIKE ?", "%"+title+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductRepositoryV1) SearchByPriceRange(minPrice, maxPrice float64) ([]models.Product, error) {
	var products []models.Product
	err := ps.DB.Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryV1) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := p.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepositoryV1) GetCommentsByProductId(id uint) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	if err := p.DB.Where("product_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func NewCategoryRepository() *CategoryRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &CategoryRepositoryV1{}
	}
	return &CategoryRepositoryV1{DB: db}
}

func (c *CategoryRepositoryV1) GetCategories() ([]models.Category, error) {

	categories := make([]models.Category, 0)

	if err := c.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil

}

func (c *CategoryRepositoryV1) CreateCategory(category *models.Category) error {
	return c.DB.Create(category).Error
}

func (c *CategoryRepositoryV1) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (c *CategoryRepositoryV1) UpdateCategory(category *models.Category) error {
	return c.DB.Save(category).Error
}

func (c *CategoryRepositoryV1) DeleteCategory(id uint) error {
	return c.DB.Delete(&models.Category{}, id).Error
}
