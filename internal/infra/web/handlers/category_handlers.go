package handlers

import (
	"encoding/json"
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
	categories, err := h.CategoryService.GetCategoriesList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setJsonContentType(w)
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, errorEncodingResponse, http.StatusInternalServerError)
		return
	}
}
