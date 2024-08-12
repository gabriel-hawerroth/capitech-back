package services

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"

type CategoryService struct {
	CategoryRepository repositories.CategoryRepository
}

func NewCategoryService(CategoryRepository repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: CategoryRepository,
	}
}
