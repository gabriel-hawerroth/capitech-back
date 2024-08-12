package services

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"

type ShoppingCartService struct {
	ShoppingCartRepository repositories.ShoppingCartRepository
}

func NewShoppingCartService(ShoppingCartRepository repositories.ShoppingCartRepository) *ShoppingCartService {
	return &ShoppingCartService{
		ShoppingCartRepository: ShoppingCartRepository,
	}
}
