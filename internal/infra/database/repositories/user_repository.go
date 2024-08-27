package repositories

import (
	"database/sql"

	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/entity"
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

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	row := r.DB.QueryRow("SELECT * FROM users WHERE email = $1", email)

	var user entity.User
	if err := scanUser(row, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateNewUser(dto dto.CreateUserDTO) error {
	_, err := r.DB.Exec("INSERT INTO users (email, password, active) VALUES ($1, $2, true)", dto.Email, dto.Password)
	return err
}

func scanUser(row *sql.Row, user *entity.User) error {
	return row.Scan(
		&user.Id, &user.Email, &user.Password,
		&user.Active, &user.CreatedAt,
	)
}
