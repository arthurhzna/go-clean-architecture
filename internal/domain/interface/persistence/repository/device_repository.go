package repository

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
)

type DeviceRepository interface {
	Create(
		ctx context.Context,
		device *entity.Device,
	) error

	FindByID(
		ctx context.Context,
		id int64,
	) (*entity.Device, error)

	Update(
		ctx context.Context,
		device *entity.Device,
	) error

	Delete(
		ctx context.Context,
		id int64,
	) error
}
