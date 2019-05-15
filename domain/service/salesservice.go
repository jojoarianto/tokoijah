package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

// SalesService contract
type SalesService interface {
	Add(model.Sales) error
	GetAll() ([]model.Sales, error)
}
