package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

type ProductService interface {
	Add(model.Product) error
	GetAll() ([]model.Product, error)
}
