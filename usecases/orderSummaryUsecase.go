package usecases

import (
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/repositories"
)

type IOrderSummaryUseCase interface {
	GetDailyReportService() (*models.OrderSummary, error)
	GetMonthlyReportService() (*models.OrderSummary, error)
	GetAllReportService() (*models.OrderSummary, error)
}

type OrderSummaryUseCase struct {
	repo repositories.IOrderRepository
}

func NewOrderSummaryUseCase(repo repositories.IOrderRepository) IOrderSummaryUseCase {
	return &OrderSummaryUseCase{repo}
}

func (p *OrderSummaryUseCase) GetDailyReportService() (*models.OrderSummary, error) {
	return p.repo.FindDailyReportOder()
}

func (p *OrderSummaryUseCase) GetMonthlyReportService() (*models.OrderSummary, error) {
	return p.repo.FindMonthlyReportOder()
}

func (p *OrderSummaryUseCase) GetAllReportService() (*models.OrderSummary, error) {
	return p.repo.FindReportOrder()
}
