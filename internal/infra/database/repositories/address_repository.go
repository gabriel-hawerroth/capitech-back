package repositories

import "database/sql"

type AddressRepository struct {
	DB *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}
