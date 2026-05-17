package repository

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
)

type DeviceLogRepository interface {
	Create(ctx context.Context, log *entity.DeviceLog) error
}
