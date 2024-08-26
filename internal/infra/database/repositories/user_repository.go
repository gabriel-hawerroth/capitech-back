package repositories

import (
	"database/sql"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	query := `
		SELECT 1
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var exists bool
	err := r.DB.QueryRow(query, email).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *UserRepository) CreateNewUser(dto dto.CreateUserDTO) error {
	_, err := r.DB.Exec("INSERT INTO users (email, password, active) VALUES ($1, $2, true)", dto.Email, dto.Password)
	return err
}
