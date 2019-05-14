package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// StockOut struct for stockout
type StockOut struct {
	gorm.Model
	StockOutTime  time.Time `validate:"required" json:"StockOutTime"`
	SalesID       int       `json:"SalesID" sql:"type:integer REFERENCES sales(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	ProductID     int       `validate:"required,numeric,min=1" json:"ProductID" sql:"type:integer REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE"`
	Product       Product   `json:"Product"`
	Qty           int       `validate:"required,numeric" json:"Qty"`
	SellPrice     int       `validate:"omitempty,numeric" json:"SellPrice"`
	TotalPrice    int       `json:"TotalPrice"`
	StatusOutCode int       `validate:"required,numeric,min=1" json:"StatusOutCode"` // 1. Terjual, 2. Barang Hilang, 3. Barang Rusak, 4 Barang Sample
}
