package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type stockOutService struct {
	stockOutRepo repository.StockOutRepo
}

// NewStockOutService method service init
func NewStockOutService(stockoutrepo repository.StockOutRepo) *stockOutService {
	return &stockOutService{
		stockOutRepo: stockoutrepo,
	}
}

// Add service to create new purchase
func (so *stockOutService) Add(stockout model.StockOut) error {
	err := so.stockOutRepo.Add(stockout)
	if err != nil {
		return err
	}

	return nil
}
