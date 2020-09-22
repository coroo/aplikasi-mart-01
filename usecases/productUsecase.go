package usecases

import (
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/repositories"
)

type IProductUseCase interface {
	GetProductByIdService(id string) (*models.Product, error)
	DeleteProductByIdService(id string) (*models.Product, error)
	RegisterNewProductService(product models.Product) (*models.Product, error)
}

type ProductUseCase struct {
	repo repositories.IProductRepository
}

func NewProductUseCase(repo repositories.IProductRepository) IProductUseCase {
	return &ProductUseCase{repo}
}

func (p *ProductUseCase) RegisterNewProductService(product models.Product) (*models.Product, error) {
	return p.repo.Insert(product)
}

func (p *ProductUseCase) GetProductByIdService(id string) (*models.Product, error) {
	return p.repo.FindOneById(id)
}

func (p *ProductUseCase) DeleteProductByIdService(id string) (*models.Product, error) {
	return p.repo.DeleteOneById(id)
}
