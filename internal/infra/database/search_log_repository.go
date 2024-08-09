package database

import "database/sql"

type SearchLogRepository struct {
	DB *sql.DB
}

func NewSearchLogRepository(db *sql.DB) *SearchLogRepository {
	return &SearchLogRepository{DB: db}
}
