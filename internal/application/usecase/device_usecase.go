package usecase

import (
	"context"

	apprequest "github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	appresponse "github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/entity"

	repositoryinterface "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"
)

type DeviceUseCaseInterface interface {
	Create(
		ctx context.Context,
		req *apprequest.CreateDeviceRequest,
	) (*appresponse.CreateDeviceResponse, error)
}

type DeviceUseCase struct {
	deviceRepo repositoryinterface.DeviceRepository
}

func NewDeviceUseCase(
	deviceRepo repositoryinterface.DeviceRepository,
) DeviceUseCaseInterface {
	return &DeviceUseCase{
		deviceRepo: deviceRepo,
	}
}

func (u *DeviceUseCase) Create(
	ctx context.Context,
	req *apprequest.CreateDeviceRequest,
) (*appresponse.CreateDeviceResponse, error) {

	device := &entity.Device{
		Name: req.Name,
	}

	err := u.deviceRepo.Create(ctx, device)
	if err != nil {
		return nil, err
	}

	return &appresponse.CreateDeviceResponse{
		Device: appresponse.DeviceResponse{
			ID:   device.ID,
			Name: device.Name,
		},
	}, nil
}
