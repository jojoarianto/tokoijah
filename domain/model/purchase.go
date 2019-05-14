package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Purchase struct for purchase (pembelian barang barang)
type Purchase struct {
	gorm.Model
	PurchaseTime time.Time `json:"PurchaseTime"`
	ProductID    int       `json:"ProductID" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Product      Product   `json:"Product"`
	OrderQty     int       `validate:"required,min=1" json:"OrderQty"`
	ReceivedQty  int       `json:"ReceivedQty"`
	Price        int       `validate:"required,min=1" json:"Price"`
	TotalPrice   int       `json:"TotalPrice"`
	Receipt      string    `json:"Receipt"`
	Progress     []StockIn `gorm:"ForeignKey:PurchaseID" json:"StockIn"`
	StausInCode  int       `json:"PurchaseCode"`
}
