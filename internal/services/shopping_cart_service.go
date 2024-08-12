package services

import "github.com/gabriel-hawerroth/capitech-back/internal/infra/database"

type ShoppingCartService struct {
	ShoppingCartRepository database.ShoppingCartRepository
}

func NewShoppingCartService(ShoppingCartRepository database.ShoppingCartRepository) *ShoppingCartService {
	return &ShoppingCartService{
		ShoppingCartRepository: ShoppingCartRepository,
	}
}
