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

func (s *ProductService) GetFilteredProducts(params dto.ProductQueryParams) ([]*entity.Product, error) {
	if params.Pagination.Size >= 50 {
		params.Pagination.Size = 50
	}

	return s.ProductRepository.GetFilteredProducts(params)
}

func (s *ProductService) Create(dto dto.CreateProductDto) (*entity.Product, error) {
	return s.ProductRepository.Create(dto)
}
