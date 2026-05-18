package persistence

import (
	"context"

	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/interface/persistence/repository"
)

type UnitOfWork interface {
	WithTransaction(
		ctx context.Context,
		fn func(UnitOfWork) error,
	) error

	DeviceRepository() repositoryiface.DeviceRepository
}
