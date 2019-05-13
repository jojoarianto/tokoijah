package repository

import "github.com/jojoarianto/tokoijah/domain/model"

type PurchaseRepo interface {
	Add(model.Purchase) error
}
