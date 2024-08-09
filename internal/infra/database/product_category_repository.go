package database

import "database/sql"

type ProductCategoryRepository struct {
	DB *sql.DB
}

func NewProductCategoryRepository(db *sql.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{DB: db}
}
