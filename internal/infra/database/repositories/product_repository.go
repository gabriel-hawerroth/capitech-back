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

func (r *ProductRepository) GetById(id *int) (*entity.Product, error) {
	row := r.DB.QueryRow("SELECT * FROM product WHERE id = $1", *id)

	var product entity.Product
	if err := scanProduct(row, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Create(product dto.SaveProductDTO) (*entity.Product, error) {
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

func (r *ProductRepository) Update(id int, product dto.SaveProductDTO) (*entity.Product, error) {
	query := `
		UPDATE product
		SET name = $2, description = $3, price = $4, category_id = $5, stock_quantity = $6
		WHERE id = $1
		RETURNING *
	`

	row := r.DB.QueryRow(query, id,
		product.Name, product.Description, product.Price, product.CategoryId, product.StockQuantity,
	)

	var updatedProduct entity.Product
	err := scanProduct(row, &updatedProduct)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
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

func (r *ProductRepository) ChangeImage(productId *int, image *string) error {
	_, err := r.DB.Exec("UPDATE product SET image = $2 WHERE id = $1", productId, image)
	return err
}

func (r *ProductRepository) RemoveImage(productId *int) error {
	_, err := r.DB.Exec("UPDATE product SET image = NULL WHERE id = $1", productId)
	return err
}

func (r *ProductRepository) ChangePrice(productId int, newPrice float64) error {
	_, err := r.DB.Exec("UPDATE product SET price = $2 WHERE id = $1", productId, newPrice)
	return err
}

func (r *ProductRepository) ChangeStockQuantity(productId int, newStockQuantity int) error {
	_, err := r.DB.Exec("UPDATE product SET stock_quantity = $2 WHERE id = $1", productId, newStockQuantity)
	return err
}

func (r *ProductRepository) GetTrendingProductsList() ([]*dto.HomeProductDTO, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.price,
			p.image,
			COUNT(1) AS totalSearchs
		FROM
			search_log sl
			JOIN product p ON (
				(sl.field_key = 'id' AND sl.field_value::int = p.id)
				OR (sl.field_key = 'name' AND p.name like '%' || sl.field_value || '%')
				OR (sl.field_key = 'category' AND p.category_id = sl.field_value::int)
			)
		GROUP BY
			p.id,
			p.name,
			p.price,
			p.image
		ORDER BY
			totalSearchs DESC
		LIMIT 12
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*dto.HomeProductDTO, 0)
	for rows.Next() {
		product := &dto.HomeProductDTO{}
		if err := scanHomeProducts(rows, product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetBestSellingProductsList() ([]*dto.HomeProductDTO, error) {
	query := `
		SELECT
			prd.id,
			prd.name,
			prd.price,
			prd.image,
			SUM(pi.quantity) AS totalSales
		FROM
			purchase p
			JOIN purchase_item pi ON p.id = pi.purchase_id
			JOIN product prd ON pi.product_id = prd.id
		GROUP BY
			prd.id,
			prd.name,
			prd.price,
			prd.image
		ORDER BY
			totalSales DESC
		LIMIT 12
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*dto.HomeProductDTO, 0)
	for rows.Next() {
		product := &dto.HomeProductDTO{}
		if err := scanHomeProducts(rows, product); err != nil {
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

func replacePlaceholders(query string, numArgs int) string {
	for i := 1; i <= numArgs; i++ {
		placeholder := fmt.Sprintf("$%d", i)
		query = strings.Replace(query, "?", placeholder, 1)
	}
	return query
}

func scanHomeProducts(rows *sql.Rows, product *dto.HomeProductDTO) error {
	var totalSearchs *int
	return rows.Scan(&product.Id, &product.Name, &product.Price, &product.Image, &totalSearchs)
}
