package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type StockOutRepo interface {
	Add(model.StockOut) error
}
