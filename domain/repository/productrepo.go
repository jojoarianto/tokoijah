package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type ProductRepo interface {
	GetAll() ([]model.Product, error)
}
