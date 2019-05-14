package model

import (
	"github.com/jinzhu/gorm"
)

// Product is data structure for product entity
type Product struct {
	gorm.Model
	Sku    string `validate:"required" gorm:"type:varchar(100);unique_index" json:"Sku"`
	Name   string `validate:"required" json:"Name"`
	Stocks int    `json:"Stocks"`
}
