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

func (s *ProductService) GetFilteredProducts(params dto.ProductQueryParams) (*dto.PaginationResponse[entity.Product], error) {
	content, err := s.ProductRepository.GetFilteredProducts(params)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.ProductRepository.GetFilteredProductsCount(params)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[entity.Product]{
		Content:    content,
		TotalItems: totalItems,
	}, nil
}

func (s *ProductService) Create(dto dto.CreateProductDto) (*entity.Product, error) {
	return s.ProductRepository.Create(dto)
}
