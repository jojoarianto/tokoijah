package sqlite3

import (
	"fmt"

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

// GetByID method to retrieve a single data product by id
func (pr *productRepo) GetByID(productID int) (model.Product, error) {
	product := model.Product{}
	if err := pr.Conn.First(&product, productID).Error; err != nil {
		return product, err
	}
	return product, nil
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

// AddMany method to add many new products
func (pr *productRepo) AddMany(products []model.Product) error {
	var sqlStr string
	for _, product := range products {
		sqlStr += fmt.Sprintf("INSERT INTO products (created_at, updated_at, sku, name, stocks) VALUES (datetime('now','localtime'), datetime('now','localtime'),'%s', '%s', %d); ", product.Sku, product.Name, product.Stocks)
	}
	pr.Conn.Exec(sqlStr)
	return nil
}
