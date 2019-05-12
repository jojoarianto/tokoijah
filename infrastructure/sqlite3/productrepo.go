package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type productRepo struct {
	Conn *gorm.DB
}

// NewProductRepo method repo init
func NewProductRepo(conn *gorm.DB) repository.ProductRepo {
	return &productRepo{Conn: conn}
}

// GetAll method to get all data product
func (pr *productRepo) GetAll() ([]model.Product, error) {
	products := []model.Product{}
	if err := pr.Conn.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

// Add method to add new product
func (pr *productRepo) Add(product model.Product) error {
	if err := pr.Conn.Save(&product).Error; err != nil {
		return err
	}
	return nil
}
