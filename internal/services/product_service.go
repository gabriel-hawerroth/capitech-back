package services

import (
	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/entity"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s *ProductService) Create(dto dto.CreateProductDto) (*entity.Product, error) {
	return s.ProductRepository.Create(dto)
}
