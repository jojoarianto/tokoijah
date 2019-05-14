package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Sales struct for sales
type Sales struct {
	gorm.Model
	SalesTime       time.Time  `validate:"required" json:"SalesTime"`
	TransactionCode string     `json:"TransactionCode"`
	Items           []StockOut `json:"Items"`
}
