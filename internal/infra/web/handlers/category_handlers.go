package handlers

import (
	"net/http"

	"github.com/gabriel-hawerroth/capitech-back/internal/services"
)

type CategoryHandler struct {
	CategoryService services.CategoryService
}

func NewCategoryHandler(CategoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: CategoryService,
	}
}

func (h *CategoryHandler) GetCategoriesList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
