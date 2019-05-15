package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type SalesRepo interface {
	Add(model.Sales) error
	GetAll() ([]model.Sales, error)
}
