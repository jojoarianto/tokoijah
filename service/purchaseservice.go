package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type purchaseService struct {
	purchaseRepo repository.PurchaseRepo
	productRepo  repository.ProductRepo
}

// NewPurchaseService method service init
func NewPurchaseService(purchaserepo repository.PurchaseRepo, productRepo repository.ProductRepo) *purchaseService {
	return &purchaseService{
		purchaseRepo: purchaserepo,
		productRepo:  productRepo,
	}
}

// Add service to create new purchase
func (ps *purchaseService) Add(purchase model.Purchase) (model.Purchase, error) {
	product, err := ps.productRepo.GetByID(purchase.ProductID)
	if err != nil {
		return purchase, nil
	}
	purchase.Product = product
	purchase.TotalPrice = purchase.Price * purchase.OrderQty

	purchase, err = ps.purchaseRepo.Add(purchase)
	if err != nil {
		return purchase, nil
	}
	return purchase, nil
}

// GetByID service to retrieve a data product by id
func (ps *purchaseService) GetByID(purchaseID int) (model.Purchase, error) {
	purchase, err := ps.purchaseRepo.GetByID(purchaseID)
	if err != nil {
		return purchase, err
	}
	return purchase, nil
}

// GetAll service to retrieve all data purchase
func (ps *purchaseService) GetAll() ([]model.Purchase, error) {
	purchases, err := ps.purchaseRepo.GetAll()
	if err != nil {
		return purchases, err
	}
	return purchases, nil
}
