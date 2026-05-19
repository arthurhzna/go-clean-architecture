package repository

import (
	"database/sql"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
	domainrepo "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domainrepo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *entity.User) error {
	query := `
		INSERT INTO users (
			uuid,
			name,
			email,
			password,
			role_id,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		user.UUID.String(),
		user.Name,
		user.Email,
		user.Password,
		user.RoleID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *userRepository) FindByID(id string) (*entity.User, error) {
	query := `
		SELECT
			id,
			uuid,
			name,
			email,
			password,
			role_id,
			created_at,
			updated_at
		FROM users
		WHERE uuid = ?
	`

	var user entity.User
	var uuidStr string

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&uuidStr,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	query := `
		SELECT
			id,
			uuid,
			name,
			email,
			password,
			role_id,
			created_at,
			updated_at
		FROM users
		WHERE email = ?
	`

	var user entity.User
	var uuidStr string

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&uuidStr,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	query := `
		UPDATE users
		SET
			name = ?,
			email = ?,
			password = ?,
			role_id = ?,
			updated_at = ?
		WHERE uuid = ?
	`

	_, err := r.db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.RoleID,
		user.UpdatedAt,
		user.UUID.String(),
	)

	return err
}

func (r *userRepository) Delete(id string) error {
	query := `
		DELETE FROM users
		WHERE uuid = ?
	`

	_, err := r.db.Exec(query, id)

	return err
}
