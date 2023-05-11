package service

type OrderServiceInterface interface {
	CreateOrder(order *Order) error
	UpdateOrder(order *Order) error
	DeleteOrder(id uint) error
	GetOrders(userId uint) ([]Order, error)
	GetOrder(id uint) *Order
}

type OrderServiceV1 struct {
	orderRepos OrderReposInterface
}

func NewOrderService() OrderServiceInterface {
	return OrderServiceV1{orderRepos: NewOrderRepo()}
}

func (o OrderServiceV1) CreateOrder(order *Order) error {
	return o.orderRepos.CreateOrder(order)
}

func (o OrderServiceV1) UpdateOrder(order *Order) error {
	return o.orderRepos.UpdateOrder(order)
}

func (o OrderServiceV1) DeleteOrder(id uint) error {
	return o.orderRepos.DeleteOrder(id)
}

func (o OrderServiceV1) GetOrders(userId uint) ([]Order, error) {
	return o.orderRepos.GetOrders(userId)
}

func (o OrderServiceV1) GetOrder(id uint) *Order {
	return o.orderRepos.GetOrder(id)
}
