package deliveries

import (
	"fmt"
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/usecases"
)

type IOrderDelivery interface {
	GetOrderById(id string) (*models.Order, error)
	DeleteOrderById(id string) (*models.Order, error)
	PrintOneOrder(result *models.Order)
	RegisterNewOrder(orderWithProduct models.OrderWithProduct) (*models.OrderWithProduct, error)
}

type OrderDelivery struct {
	orderService usecases.IOrderUseCase
}

func NewOrderDelivery(service usecases.IOrderUseCase) IOrderDelivery {
	return &OrderDelivery{service}
}

func (pd *OrderDelivery) RegisterNewOrder(orderWithProduct models.OrderWithProduct) (*models.OrderWithProduct, error) {
	var err error
	err = orderWithProduct.Order.Validate()
	if err != nil {
		return nil, err
	}
	return pd.orderService.RegisterNewOrderService(orderWithProduct)
}

func (pd *OrderDelivery) GetOrderById(id string) (*models.Order, error) {
	return pd.orderService.GetOrderByIdService(id)
}

func (pd *OrderDelivery) DeleteOrderById(id string) (*models.Order, error) {
	return pd.orderService.DeleteOrderByIdService(id)
}

func (pd *OrderDelivery) PrintOneOrder(result *models.Order) {
	fmt.Printf("%v\n", result)
}
