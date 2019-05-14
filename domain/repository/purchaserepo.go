package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type PurchaseRepo interface {
	Add(model.Purchase) (model.Purchase, error)
	GetAll() ([]model.Purchase, error)
	GetByID(purchaseID int) (model.Purchase, error)
}
