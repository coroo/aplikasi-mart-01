package deliveries

import (
	"fmt"
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/usecases"
)

type IOrderSummaryDelivery interface {
	GetDailyReport() (*models.OrderSummary, error)
	GetMonthlyReport() (*models.OrderSummary, error)
	GetAllReport() (*models.OrderSummary, error)
	PrintOneOrder(result *models.OrderSummary)
}

type OrderSummaryDelivery struct {
	orderSummaryService usecases.IOrderSummaryUseCase
}

func NewOrderSummaryDelivery(service usecases.IOrderSummaryUseCase) IOrderSummaryDelivery {
	return &OrderSummaryDelivery{service}
}

func (pd *OrderSummaryDelivery) GetDailyReport() (*models.OrderSummary, error) {
	return pd.orderSummaryService.GetDailyReportService()
}

func (pd *OrderSummaryDelivery) GetMonthlyReport() (*models.OrderSummary, error) {
	return pd.orderSummaryService.GetMonthlyReportService()
}

func (pd *OrderSummaryDelivery) GetAllReport() (*models.OrderSummary, error) {
	return pd.orderSummaryService.GetAllReportService()
}

func (pd *OrderSummaryDelivery) PrintOneOrder(result *models.OrderSummary) {
	fmt.Println("\nOrder Quantity:", result.OrderCount)
	fmt.Println("Order Total Amount:", result.OrderTotal)
}
