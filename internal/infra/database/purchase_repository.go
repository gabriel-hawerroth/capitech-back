package database

import "database/sql"

type PurchaseRepository struct {
	DB *sql.DB
}

func NewPurchaseRepository(db *sql.DB) *PurchaseRepository {
	return &PurchaseRepository{DB: db}
}
