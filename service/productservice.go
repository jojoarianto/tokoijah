package service

import (
	"github.com/jojoarianto/tokoijah/domain/model"
	"github.com/jojoarianto/tokoijah/domain/repository"
)

type productService struct {
	productRepo repository.ProductRepo
}

// NewProductService method service init
func NewProductService(productRepo repository.ProductRepo) *productService {
	return &productService{
		productRepo: productRepo,
	}
}

// GetAll service to retrieve all data product
func (ps *productService) GetAll() ([]model.Product, error) {
	products, err := ps.productRepo.GetAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

// GetByID service to retrieve a data product by id
func (ps *productService) GetByID(productID int) (model.Product, error) {
	product, err := ps.productRepo.GetByID(productID)
	if err != nil {
		return product, err
	}
	return product, nil
}

// Add service to create new product
func (ps *productService) Add(product model.Product) error {
	err := ps.productRepo.Add(product)
	if err != nil {
		return err
	}
	return nil
}

// AddMany service to insert many product to db
func (ps *productService) AddMany(products []model.Product) error {
	err := ps.productRepo.AddMany(products)
	if err != nil {
		return err
	}
	return nil
}
