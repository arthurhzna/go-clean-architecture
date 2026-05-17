package usecase

import (
	"context"

	apprequest "github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	appresponse "github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"

	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/interface/persistence/repository"
)

type DeviceUseCase struct {
	deviceRepo repositoryiface.DeviceRepository
}

func NewDeviceUseCase(
	deviceRepo repositoryiface.DeviceRepository,
) *DeviceUseCase {
	return &DeviceUseCase{
		deviceRepo: deviceRepo,
	}
}

func (u *DeviceUseCase) FindByID(
	ctx context.Context,
	req *apprequest.FindDeviceByIDRequest,
) (*appresponse.DeviceResponse, error) {

	device, err := u.deviceRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &appresponse.DeviceResponse{
		ID:   device.ID,
		Name: device.Name,
	}, nil
}
