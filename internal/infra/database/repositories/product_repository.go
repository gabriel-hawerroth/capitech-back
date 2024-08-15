package repositories

import (
	"database/sql"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product dto.CreateProductDto) (*entity.Product, error) {
	query := `
		INSERT INTO product (name, description, price, category_id, stock_quantity)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`

	row := r.DB.QueryRow(query,
		product.Name, product.Description, product.Price, product.CategoryId, product.StockQuantity,
	)

	var createdProduct entity.Product
	err := scanProduct(row, &createdProduct)
	if err != nil {
		return nil, err
	}

	return &createdProduct, nil
}

func scanProduct(row *sql.Row, product *entity.Product) error {
	return row.Scan(
		&product.Id, &product.Name, &product.Description, &product.Price,
		&product.CategoryId, &product.StockQuantity, &product.Image,
	)
}
