package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type purchaseService struct {
	purchaseRepo repository.PurchaseRepo
}

// NewPurchaseService method service init
func NewPurchaseService(purchaserepo repository.PurchaseRepo) *purchaseService {
	return &purchaseService{
		purchaseRepo: purchaserepo,
	}
}

// Add service to create new purchase
func (ps *purchaseService) Add(purchase model.Purchase) error {
	err := ps.purchaseRepo.Add(purchase)
	if err != nil {
		return err
	}
	return nil
}
