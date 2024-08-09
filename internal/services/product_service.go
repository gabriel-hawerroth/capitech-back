package services

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database"

type ProductService struct {
	ProductRepository database.ProductRepository
}

func NewProductService(productRepository database.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}
