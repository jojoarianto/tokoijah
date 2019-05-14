package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Purchase struct for purchase (pembelian barang barang)
type Purchase struct {
	gorm.Model
	PurchaseTime time.Time `json:"purchase_time"`
	ProductID    int       `json:"product_id" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Product      Product   `json:"product"`
	OrderQty     int       `validate:"required,min=1" json:"order_qty"`
	ReceivedQty  int       `json:"received_qty"`
	Price        int       `validate:"required,min=1" json:"price"`
	TotalPrice   int       `json:"total_price"`
	Receipt      string    `json:"receipt"`
	Progress     []StockIn `gorm:"ForeignKey:PurchaseID" json:"stock_in"`
	StausInCode  int       `json:"purchase_code"`
}
