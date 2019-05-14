package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type purchaseRepo struct {
	Conn *gorm.DB
}

// NewPurchaseRepo method repo init
func NewPurchaseRepo(conn *gorm.DB) repository.PurchaseRepo {
	return &purchaseRepo{Conn: conn}
}

// Add method to add new purchase
func (pr *purchaseRepo) Add(purchase model.Purchase) (model.Purchase, error) {
	if err := pr.Conn.Save(&purchase).Error; err != nil {
		return purchase, err
	}
	return purchase, nil
}

// GetByID method to retrieve a single data purchase by id
func (pr *purchaseRepo) GetByID(purchaseID int) (model.Purchase, error) {
	purchase := model.Purchase{}
	if err := pr.Conn.First(&purchase, purchaseID).Error; err != nil {
		return purchase, err
	}
	return purchase, nil
}
