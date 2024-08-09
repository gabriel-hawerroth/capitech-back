package entity

type PurchaseItem struct {
	Id           int     `json:"id"`
	PurchaseId   int     `json:"purchase_id"`
	ProductId    int     `json:"product_id"`
	ProductValue float64 `json:"product_value"`
	Quantity     int     `json:"quantity"`
}
