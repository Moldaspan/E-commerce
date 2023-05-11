package service

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/repositories"
)

type OrderServiceInterface interface {
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
	GetOrders(userId uint) ([]models.Order, error)
	GetOrder(id uint) *models.Order
}

type OrderServiceV1 struct {
	orderRepos repositories.OrderReposInterface
}

func NewOrderService() OrderServiceInterface {
	return OrderServiceV1{orderRepos: repositories.NewOrderRepo()}
}

func (o OrderServiceV1) CreateOrder(order *models.Order) error {
	return o.orderRepos.CreateOrder(order)
}

func (o OrderServiceV1) UpdateOrder(order *models.Order) error {
	return o.orderRepos.UpdateOrder(order)
}

func (o OrderServiceV1) DeleteOrder(id uint) error {
	return o.orderRepos.DeleteOrder(id)
}

func (o OrderServiceV1) GetOrders(userId uint) ([]models.Order, error) {
	return o.orderRepos.GetOrders(userId)
}

func (o OrderServiceV1) GetOrder(id uint) *models.Order {
	return o.orderRepos.GetOrder(id)
}
