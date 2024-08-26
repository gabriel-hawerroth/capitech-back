package dto

type SaveProductDTO struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CategoryId    int     `json:"category_id"`
	StockQuantity int     `json:"stock_quantity"`
}

type ProductFilter struct {
	Name       *string `json:"name,omitempty"`
	MinPrice   float64 `json:"minPrice"`
	MaxPrice   float64 `json:"maxPrice"`
	Categories []int   `json:"categories,omitempty"`
}

type ProductQueryParams struct {
	Filters    ProductFilter `json:"filters"`
	Pagination Pagination    `json:"pagination"`
}

type ChangeProductPriceDTO struct {
	NewPrice float64 `json:"newPrice"`
}

type ChangeProductStockQuantityDTO struct {
	NewStockQuantity int `json:"newStockQuantity"`
}

type HomeProductDTO struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image *string `json:"image"`
}
