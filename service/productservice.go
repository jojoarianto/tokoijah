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

	products, err := ps.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}
