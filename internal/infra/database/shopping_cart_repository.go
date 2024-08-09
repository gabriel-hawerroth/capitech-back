package database

import "database/sql"

type ShoppingCartRepository struct {
	DB *sql.DB
}

func NewShoppingCartRepository(db *sql.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{DB: db}
}
