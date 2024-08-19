package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
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

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, errorCastingParams, http.StatusBadRequest)
		return
	}

	product, err := h.ProductService.GetById(productId)
	if err != nil {
		setHttpError(w, err, http.StatusBadRequest)
		return
	}

	setJsonContentType(w)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, errorEncodingResponse, http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.SaveProductDto
	err := json.NewDecoder(r.Body).Decode(&data)
	checkDecodeError(err, w)

	product, err := h.ProductService.Create(data)
	if err != nil {
		setHttpError(w, err, http.StatusInternalServerError)
		return
	}

	setJsonContentType(w)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, errorEncodingResponse, http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, errorCastingParams, http.StatusBadRequest)
		return
	}

	var data dto.SaveProductDto
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, errorDecodingRequestBody, http.StatusInternalServerError)
		return
	}

	product, err := h.ProductService.Update(productId, data)
	if err != nil {
		setHttpError(w, err, http.StatusInternalServerError)
		return
	}

	setJsonContentType(w)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, errorEncodingResponse, http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) GetProductsList(w http.ResponseWriter, r *http.Request) {
	filters, pagination, err := parseListQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queryParams := dto.ProductQueryParams{
		Filters:    filters,
		Pagination: pagination,
	}

	products, err := h.ProductService.GetFilteredProducts(queryParams)
	if err != nil {
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}

	setJsonContentType(w)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, errorEncodingResponse, http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) GetTrendingProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ProductHandler) GetBestSellingProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ProductHandler) GetUserSearchHistory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ProductHandler) ChangePrice(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ProductHandler) ChangeStockQuantity(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ProductHandler) ChangeImage(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, errorCastingParams, http.StatusBadRequest)
		return
	}

	err = h.ProductService.ChangeImage(productId, w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload image: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) RemoveImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *ProductHandler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func parseListQueryParams(r *http.Request) (dto.ProductFilter, dto.Pagination, error) {
	filtersParam := r.URL.Query().Get("filters")
	paginationParam := r.URL.Query().Get("pagination")

	var filters dto.ProductFilter
	var pagination dto.Pagination

	if filtersParam != "" {
		if err := json.Unmarshal([]byte(filtersParam), &filters); err != nil {
			return filters, pagination, fmt.Errorf("invalid filters parameter: %v", err)
		}
	}

	if paginationParam != "" {
		if err := json.Unmarshal([]byte(paginationParam), &pagination); err != nil {
			return filters, pagination, fmt.Errorf("invalid pagination parameter: %v", err)
		}
	}

	return filters, pagination, nil
}
