package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// StockIn struct for stock_in (data barang masuk)
type StockIn struct {
	gorm.Model
	StockInTime time.Time `validate:"required" json:"StockInTime"`
	PurchaseID  int       `validate:"required,numeric,min=1" json:"PurchaseID" sql:"type:integer REFERENCES purchase(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Qty         int       `validate:"required,numeric,min=1" json:"Qty"`
}
