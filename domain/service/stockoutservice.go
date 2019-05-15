package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

// StockOutService contract
type StockOutService interface {
	Add(model.StockOut) error
}
