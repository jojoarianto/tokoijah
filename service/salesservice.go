package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type salesService struct {
	salesRepo repository.SalesRepo
}

// NewSalesService method service init
func NewSalesService(salesrepo repository.SalesRepo) *salesService {
	return &salesService{
		salesRepo: salesrepo,
	}
}

// Add service to create new sales
func (ps *salesService) Add(sales model.Sales) error {
	err := ps.salesRepo.Add(sales)
	if err != nil {
		return nil
	}
	return nil
}
