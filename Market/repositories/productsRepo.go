package repositories

import (
	"github.com/Moldaspan/E-commerce/comments"
	"github.com/Moldaspan/E-commerce/settings"
	"gorm.io/gorm"
	"log"
)

type ProductRepositoryInterface interface {
	CreateProduct(*Product) error
	GetProductByID(uint) (*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(uint) error
	GetProducts() ([]Product, error)
	GetCommentsByProductId(uint) ([]*comments.Comment, error)
	GetProductAverageRating(uint) (float32, error)

	SearchByName(name string) ([]Product, error)
	SearchByPriceRange(minPrice, maxPrice float64) ([]Product, error)
}
type CategoryRepositoryInterface interface {
	CreateCategory(category *Category) error
	GetCategoryByID(uint) (*Category, error)
	UpdateCategory(*Category) error
	DeleteCategory(uint) error
	GetCategories() ([]Category, error)
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

func (p *ProductRepositoryV1) CreateProduct(product *Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepositoryV1) UpdateProduct(product *Product) error {
	return p.DB.Save(product).Error
}

func (p *ProductRepositoryV1) DeleteProduct(id uint) error {
	return p.DB.Delete(&Product{}, id).Error
}

func (p *ProductRepositoryV1) GetProducts() ([]Product, error) {
	products := make([]Product, 0)

	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (ps ProductRepositoryV1) SearchByName(title string) ([]Product, error) {
	var products []Product
	err := ps.DB.Where("title LIKE ?", "%"+title+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductRepositoryV1) SearchByPriceRange(minPrice, maxPrice float64) ([]Product, error) {
	var products []Product
	err := ps.DB.Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryV1) GetProductByID(id uint) (*Product, error) {
	var product Product
	if err := p.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepositoryV1) GetCommentsByProductId(id uint) ([]*comments.Comment, error) {
	comments := make([]*comments.Comment, 0)
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

func (c *CategoryRepositoryV1) GetCategories() ([]Category, error) {

	categories := make([]Category, 0)

	if err := c.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil

}

func (c *CategoryRepositoryV1) CreateCategory(category *Category) error {
	return c.DB.Create(category).Error
}

func (c *CategoryRepositoryV1) GetCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (c *CategoryRepositoryV1) UpdateCategory(category *Category) error {
	return c.DB.Save(category).Error
}

func (c *CategoryRepositoryV1) DeleteCategory(id uint) error {
	return c.DB.Delete(&Category{}, id).Error
}
