package usecase

import (
	"context"

	apprequest "github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	appresponse "github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"

	persistenceiface "github.com/arthurhzna/go-clean-architecture/internal/domain/interface/persistence"
)

type DeviceLogUseCase struct {
	uow persistenceiface.UnitOfWork
}

func NewDeviceLogUseCase(
	uow persistenceiface.UnitOfWork,
) *DeviceLogUseCase {
	return &DeviceLogUseCase{
		uow: uow,
	}
}

func (u *DeviceLogUseCase) CreateDeviceWithLog(
	ctx context.Context,
	req *apprequest.CreateDeviceWithLogRequest,
) (*appresponse.CreateDeviceWithLogResponse, error) {

	var res *appresponse.CreateDeviceWithLogResponse

	err := u.uow.WithTransaction(
		ctx,
		func(txUow persistenceiface.UnitOfWork) error {

			deviceRepo := txUow.DeviceRepository()
			deviceLogRepo := txUow.DeviceLogRepository()

			device := &entity.Device{
				Name: req.Name,
			}

			err := deviceRepo.Create(ctx, device)
			if err != nil {
				return err
			}

			deviceLog := &entity.DeviceLog{
				DeviceID: device.ID,
				Action:   "CREATE_DEVICE",
			}

			err = deviceLogRepo.Create(ctx, deviceLog)
			if err != nil {
				return err
			}

			res = &appresponse.CreateDeviceWithLogResponse{
				ID:   device.ID,
				Name: device.Name,
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}
