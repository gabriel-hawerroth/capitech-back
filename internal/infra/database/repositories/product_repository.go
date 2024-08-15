package repositories

import (
	"database/sql"
	"strings"

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

func (r *ProductRepository) GetFilteredProducts(params dto.ProductQueryParams) ([]*entity.Product, error) {
	query := `
		SELECT *
		FROM product
		WHERE price BETWEEN ? AND ?
	`

	args := []interface{}{params.Filters.MinPrice, params.Filters.MaxPrice}

	if params.Filters.Name != nil {
		lowerName := strings.ToLower(*params.Filters.Name)
		query += " AND LOWER(name) LIKE ?"
		args = append(args, "%"+lowerName+"%")
	}

	if len(params.Filters.Categories) > 0 {
		query += " AND category_id IN (?" + strings.Repeat(", ?", len(params.Filters.Categories)-1) + ")"
		for _, categoryId := range params.Filters.Categories {
			args = append(args, categoryId)
		}
	}

	query += " ORDER BY id LIMIT ? OFFSET ?"
	offset := (params.Pagination.Page - 1) * params.Pagination.Size
	args = append(args, params.Pagination.Size, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		product := new(entity.Product)
		if err := scanProducts(rows, product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func scanProduct(row *sql.Row, product *entity.Product) error {
	return row.Scan(
		&product.Id, &product.Name, &product.Description, &product.Price,
		&product.CategoryId, &product.StockQuantity, &product.Image,
	)
}

func scanProducts(rows *sql.Rows, product *entity.Product) error {
	return rows.Scan(
		&product.Id, &product.Name, &product.Description, &product.Price,
		&product.CategoryId, &product.StockQuantity, &product.Image,
	)
}
