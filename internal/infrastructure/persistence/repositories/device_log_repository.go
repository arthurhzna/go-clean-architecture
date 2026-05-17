package repositories

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"
	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/interface/persistence/repository"

	dbtx "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/persistence/dbtx"
)

type deviceLogRepository struct {
	db dbtx.DBTX
}

func NewDeviceLogRepository(
	db dbtx.DBTX,
) repositoryiface.DeviceLogRepository {
	return &deviceLogRepository{
		db: db,
	}
}

func (r *deviceLogRepository) Create(
	ctx context.Context,
	log *entity.DeviceLog,
) error {

	query := `
		INSERT INTO device_logs (
			device_id,
			action
		)
		VALUES ($1, $2)
		RETURNING id
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		log.DeviceID,
		log.Action,
	).Scan(&log.ID)
}
