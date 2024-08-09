package handlers

import (
	"fmt"
	"net/http"

	"github.com/gabriel-hawerroth/capitech-back/internal/services"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

func (h *ProductHandler) GetProductsList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
