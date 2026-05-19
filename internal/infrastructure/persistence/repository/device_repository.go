package repository

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"

	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"

	dbtx "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/persistence/dbtx"
)

type deviceRepository struct {
	db dbtx.DBTX
}

func NewDeviceRepository(
	db dbtx.DBTX,
) repositoryiface.DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (r *deviceRepository) Create(
	ctx context.Context,
	device *entity.Device,
) error {

	query := `
		INSERT INTO devices (
			name
		)
		VALUES ($1)
		RETURNING id
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		device.Name,
	).Scan(&device.ID)
}

func (r *deviceRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entity.Device, error) {

	query := `
		SELECT
			id,
			name
		FROM devices
		WHERE id = $1
	`

	device := &entity.Device{}

	err := r.db.QueryRowxContext(
		ctx,
		query,
		id,
	).StructScan(device)

	if err != nil {
		return nil, err
	}

	return device, nil
}

func (r *deviceRepository) Update(
	ctx context.Context,
	device *entity.Device,
) error {

	query := `
		UPDATE devices
		SET
			name = $1
		WHERE id = $2
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		device.Name,
		device.ID,
	)

	return err
}

func (r *deviceRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `
		DELETE FROM devices
		WHERE id = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)

	return err
}
