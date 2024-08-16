package repositories

import (
	"database/sql"
	"fmt"
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

// Função auxiliar para construir a query e os argumentos
func buildProductQuery(params dto.ProductQueryParams, selectClause string) (string, []interface{}) {
	query := selectClause + `
		FROM product
		WHERE price BETWEEN ? AND ?
	`

	args := make([]any, 0)
	args = append(args, params.Filters.MinPrice, params.Filters.MaxPrice)

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

	return query, args
}

func (r *ProductRepository) GetFilteredProducts(params dto.ProductQueryParams) ([]*entity.Product, error) {
	if params.Pagination.Size >= 50 {
		params.Pagination.Size = 50
	}

	query, args := buildProductQuery(params, "SELECT *")

	query += " ORDER BY name asc OFFSET ? LIMIT ?"
	offset := (params.Pagination.Page) * params.Pagination.Size
	args = append(args, offset, params.Pagination.Size)

	query = replacePlaceholders(query, len(args))

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*entity.Product, 0)
	for rows.Next() {
		product := &entity.Product{}
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

func (r *ProductRepository) GetFilteredProductsCount(params dto.ProductQueryParams) (int, error) {
	query, args := buildProductQuery(params, "SELECT count(1)")

	query = replacePlaceholders(query, len(args))

	row := r.DB.QueryRow(query, args...)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
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

func replacePlaceholders(query string, numArgs int) string {
	for i := 1; i <= numArgs; i++ {
		placeholder := fmt.Sprintf("$%d", i)
		query = strings.Replace(query, "?", placeholder, 1)
	}
	return query
}
