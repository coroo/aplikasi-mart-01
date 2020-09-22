package deliveries

import (
	"fmt"
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/usecases"
)

type IProductDelivery interface {
	GetProductById(id string) (*models.Product, error)
	DeleteProductById(id string) (*models.Product, error)
	PrintOneProduct(result *models.Product)
	RegisterNewProduct(product models.Product) (*models.Product, error)
}

type ProductDelivery struct {
	productService usecases.IProductUseCase
}

func NewProductDelivery(service usecases.IProductUseCase) IProductDelivery {
	return &ProductDelivery{service}
}

func (pd *ProductDelivery) RegisterNewProduct(product models.Product) (*models.Product, error) {
	var err error
	err = product.Validate()
	if err != nil {
		return nil, err
	}
	return pd.productService.RegisterNewProductService(product)
}

func (pd *ProductDelivery) GetProductById(id string) (*models.Product, error) {
	return pd.productService.GetProductByIdService(id)
}

func (pd *ProductDelivery) DeleteProductById(id string) (*models.Product, error) {
	return pd.productService.DeleteProductByIdService(id)
}

func (pd *ProductDelivery) PrintOneProduct(result *models.Product) {
	fmt.Printf("%v\n", result)
}
