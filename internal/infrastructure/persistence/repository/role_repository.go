package repository

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"

	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"

	dbtx "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/persistence/dbtx"
)

type roleRepository struct {
	db dbtx.DBTX
}

func NewRoleRepository(
	db dbtx.DBTX,
) repositoryiface.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entity.Role, error) {

	query := `
		SELECT
			id,
			name
		FROM roles
		WHERE id = $1
	`

	var role entity.Role

	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).Scan(
		&role.ID,
		&role.Name,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindByName(
	ctx context.Context,
	name string,
) (*entity.Role, error) {

	query := `
		SELECT
			id,
			name
		FROM roles
		WHERE name = $1
	`

	var role entity.Role

	err := r.db.QueryRowxContext(
		ctx,
		query,
		name,
	).Scan(
		&role.ID,
		&role.Name,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}
