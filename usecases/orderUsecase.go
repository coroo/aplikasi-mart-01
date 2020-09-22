package usecases

import (
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/repositories"
)

type IOrderUseCase interface {
	GetOrderByIdService(id string) (*models.Order, error)
	DeleteOrderByIdService(id string) (*models.Order, error)
	RegisterNewOrderService(orderWithProduct models.OrderWithProduct) (*models.OrderWithProduct, error)
}

type OrderUseCase struct {
	repo repositories.IOrderRepository
}

func NewOrderUseCase(repo repositories.IOrderRepository) IOrderUseCase {
	return &OrderUseCase{repo}
}

func (p *OrderUseCase) RegisterNewOrderService(orderWithProduct models.OrderWithProduct) (*models.OrderWithProduct, error) {
	return p.repo.Insert(orderWithProduct)
}

func (p *OrderUseCase) GetOrderByIdService(id string) (*models.Order, error) {
	return p.repo.FindOneById(id)
}

func (p *OrderUseCase) DeleteOrderByIdService(id string) (*models.Order, error) {
	return p.repo.DeleteOneById(id)
}
