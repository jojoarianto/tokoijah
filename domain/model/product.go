package model

import (
	"github.com/jinzhu/gorm"
)

// Product is data structure for product entity
type Product struct {
	gorm.Model
	Sku    string `gorm:"type:varchar(100);unique_index" json:"sku"`
	Name   string `json:"name"`
	Stocks int    `json:"stocks"`
}
