package repositories

import (
	"database/sql"

	"github.com/gabriel-hawerroth/capitech-back/internal/entity"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetCategoriesList() ([]*entity.Category, error) {
	rows, err := r.DB.Query("SELECT * FROM category ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories = make([]*entity.Category, 8)

	i := 0
	for rows.Next() {
		category := &entity.Category{}
		if err := scanCategories(rows, category); err != nil {
			return nil, err
		}
		categories[i] = category
		i++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func scanCategories(rows *sql.Rows, category *entity.Category) error {
	return rows.Scan(&category.Id, &category.Description)
}
