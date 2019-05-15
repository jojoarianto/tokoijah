package sqlite3

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type stockOutRepo struct {
	Conn *gorm.DB
}

// NewStockOutRepo method repo init
func NewStockOutRepo(conn *gorm.DB) repository.StockOutRepo {
	return &stockOutRepo{Conn: conn}
}

// Add method to add new stockout
func (so *stockOutRepo) Add(stockout model.StockOut) error {
	tx := so.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return model.ErrTrxFail
	}

	// create stock in
	stockout.TotalPrice = stockout.SellPrice * stockout.Qty
	if err := tx.Create(&stockout).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update stock qty on product
	product := model.Product{}
	if err := tx.First(&product, stockout.ProductID).Error; err != nil {
		tx.Rollback()
		return err
	}
	product.Stock = product.Stock - stockout.Qty
	if product.Stock < 0 { // cek stock enough to checkout
		tx.Rollback()
		return model.ErrNotEnoughStock
	}
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
