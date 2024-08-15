package dto

type CreateProductDto struct {
	Name          string  `json:"name"`
	Description   *string `json:"description"`
	Price         float64 `json:"price"`
	CategoryId    int     `json:"category_id"`
	StockQuantity int     `json:"stock_quantity"`
}
