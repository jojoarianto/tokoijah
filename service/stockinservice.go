package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type stockInService struct {
	stockInRepo repository.StockInRepo
}

// NewStockInService method service init
func NewStockInService(stockinrepo repository.StockInRepo) *stockInService {
	return &stockInService{
		stockInRepo: stockinrepo,
	}
}

// Add service to create new purchase
func (ss *stockInService) Add(purchaseID int, stockin model.StockIn) error {
	stockin.PurchaseID = purchaseID
	err := ss.stockInRepo.Add(stockin)
	if err != nil {
		return err
	}

	return nil
}
