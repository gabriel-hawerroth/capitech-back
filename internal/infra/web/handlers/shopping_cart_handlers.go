package handlers

import (
	"net/http"

	"github.com/gabriel-hawerroth/capitech-back/internal/services"
)

type ShoppingCartHandler struct {
	ShoppingCartService services.ShoppingCartService
}

func NewShoppingCartHandler(ShoppingCartService services.ShoppingCartService) *ShoppingCartHandler {
	return &ShoppingCartHandler{
		ShoppingCartService: ShoppingCartService,
	}
}

func (h *ShoppingCartHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ShoppingCartHandler) GetUserShoppingCart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
