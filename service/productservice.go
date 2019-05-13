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
