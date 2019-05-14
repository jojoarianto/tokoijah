package sqlite3

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type stockInRepo struct {
	Conn *gorm.DB
}

// NewStockInRepo method repo init
func NewStockInRepo(conn *gorm.DB) repository.StockInRepo {
	return &stockInRepo{Conn: conn}
}

// Add method to add new purchase
func (sr *stockInRepo) Add(stockin model.StockIn) error {
	tx := sr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return model.ErrTrxFail
	}

	// create stock in
	if err := tx.Create(&stockin).Error; err != nil {
		tx.Rollback()
		return err
	}

	// cek purchae is available & get purchase
	purchase := model.Purchase{}
	if err := tx.Preload("Product").First(&purchase, stockin.PurchaseID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update stock qty & status on purchase
	purchase.ReceivedQty = purchase.ReceivedQty + stockin.Qty
	purchase.Product.Stock = purchase.Product.Stock + stockin.Qty
	if purchase.ReceivedQty >= purchase.OrderQty {
		purchase.StausInCode = 1
	}
	if err := tx.Save(&purchase).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
