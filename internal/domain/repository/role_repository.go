package repository

import "github.com/arthurhzna/go-clean-architecture/internal/domain/entity"

type RoleRepository interface {
	FindByID(id int64) (*entity.Role, error)

	FindByName(name string) (*entity.Role, error)
}
