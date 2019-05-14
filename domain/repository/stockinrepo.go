package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type StockInRepo interface {
	Add(model.StockIn) error
}
