package repository

import (
	"context"
)

type UnitOfWork interface {
	WithTransaction(
		ctx context.Context,
		fn func(UnitOfWork) error,
	) error

	DeviceRepository() DeviceRepository
	UserRepository() UserRepository
	RoleRepository() RoleRepository
}
