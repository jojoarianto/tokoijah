package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type salesRespo struct {
	Conn *gorm.DB
}

// NewSalesRepo method repo init
func NewSalesRepo(conn *gorm.DB) repository.SalesRepo {
	return &salesRespo{Conn: conn}
}

// Add method to add new purchase
func (sr *salesRespo) Add(sales model.Sales) error {
	if err := sr.Conn.Save(&sales).Error; err != nil {
		return err
	}
	return nil
}
