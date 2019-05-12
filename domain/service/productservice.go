package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
}
