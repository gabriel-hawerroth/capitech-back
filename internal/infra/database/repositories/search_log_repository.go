package repositories

import (
	"database/sql"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
)

type SearchLogRepository struct {
	DB *sql.DB
}

func NewSearchLogRepository(db *sql.DB) *SearchLogRepository {
	return &SearchLogRepository{DB: db}
}

func (r *SearchLogRepository) Save(dto dto.SaveSearchLogDTO) error {
	query := `
		INSERT INTO search_log (field_key, field_value)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(query, dto.FieldKey, dto.FieldValue)
	return err
}

func (r *SearchLogRepository) SaveWithUser(dto dto.SaveSearchLogWithUserDTO) error {
	query := `
		INSERT INTO search_log (user_id, field_key, field_value)
		VALUES ($1, $2, $3)
	`

	_, err := r.DB.Exec(query, dto.UserId, dto.FieldKey, dto.FieldValue)
	return err
}
