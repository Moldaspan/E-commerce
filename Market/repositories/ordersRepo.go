package repositories

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/settings"
	"gorm.io/gorm"
)

type OrderReposInterface interface {
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
	GetOrders(userId uint) ([]models.Order, error)
	GetOrder(id uint) *models.Order
}

type OrderReposV1 struct {
	DB *gorm.DB
}

func NewOrderRepo() OrderReposInterface {
	db, _ := settings.DbSetup()
	return OrderReposV1{DB: db}
}

func (o OrderReposV1) CreateOrder(order *models.Order) error {
	return o.DB.Create(order).Error
}

func (o OrderReposV1) UpdateOrder(order *models.Order) error {
	return o.DB.Save(order).Error
}

func (o OrderReposV1) DeleteOrder(id uint) error {
	return o.DB.Delete(&models.Order{}, id).Error
}

func (o OrderReposV1) GetOrders(userId uint) ([]models.Order, error) {
	var orders []models.Order
	err := o.DB.Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o OrderReposV1) GetOrder(id uint) *models.Order {
	var order models.Order
	o.DB.Where("id = ?", id).Find(&order)
	return &order
}
