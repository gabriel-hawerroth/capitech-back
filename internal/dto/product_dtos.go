package dto

type SaveProductDto struct {
	Name          string  `json:"name"`
	Description   *string `json:"description"`
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
