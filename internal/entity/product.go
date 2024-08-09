package entity

type Product struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CategoryId    int     `json:"category_id"`
	StockQuantity int     `json:"stock_quantity"`
	Image         string  `json:"image"`
}
