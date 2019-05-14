package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

// StockInService contract
type StockInService interface {
	Add(purchaseID int, stockin model.StockIn) error
}
