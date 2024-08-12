package services

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}
