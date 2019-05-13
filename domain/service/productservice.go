package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

// ProductService contract
type ProductService interface {
	Add(model.Product) error
	AddMany([]model.Product) error
	GetAll() ([]model.Product, error)
}
