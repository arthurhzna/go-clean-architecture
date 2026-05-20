package repository

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
)

type RoleRepository interface {
	FindByID(
		ctx context.Context,
		id int64,
	) (*entity.Role, error)

	FindByName(
		ctx context.Context,
		name string,
	) (*entity.Role, error)
}
