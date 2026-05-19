package repository

import (
	"database/sql"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
	domainrepo "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) domainrepo.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) FindByID(id int64) (*entity.Role, error) {
	query := `
		SELECT id, name
		FROM roles
		WHERE id = ?
	`

	var role entity.Role

	err := r.db.QueryRow(query, id).Scan(
		&role.ID,
		&role.Name,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindByName(name string) (*entity.Role, error) {
	query := `
		SELECT id, name
		FROM roles
		WHERE name = ?
	`

	var role entity.Role

	err := r.db.QueryRow(query, name).Scan(
		&role.ID,
		&role.Name,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}
