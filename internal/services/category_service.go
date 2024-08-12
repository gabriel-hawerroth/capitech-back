package services

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database"

type CategoryService struct {
	CategoryRepository database.CategoryRepository
}

func NewCategoryService(CategoryRepository database.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: CategoryRepository,
	}
}
