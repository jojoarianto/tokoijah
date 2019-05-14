package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

// PurchaseService contract
type PurchaseService interface {
	Add(model.Purchase) (model.Purchase, error)
}
