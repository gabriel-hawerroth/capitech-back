package repositories

import "database/sql"

type PurchaseItemRepository struct {
	DB *sql.DB
}

func NewPurchaseItemRepository(db *sql.DB) *PurchaseItemRepository {
	return &PurchaseItemRepository{DB: db}
}
