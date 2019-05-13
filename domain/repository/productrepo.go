package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type ProductRepo interface {
	Add(model.Product) error
	AddMany([]model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(productID int) (model.Product, error)
}
