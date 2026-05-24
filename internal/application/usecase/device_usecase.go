package usecase

import (
	"context"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/usecase"

	apprequest "github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	appresponse "github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"

	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"

	repositoryinterface "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"
)

type DeviceUseCase struct {
	deviceRepo repositoryinterface.DeviceRepository
}

func NewDeviceUseCase(
	deviceRepo repositoryinterface.DeviceRepository,
) usecase.DeviceUseCase {
	return &DeviceUseCase{
		deviceRepo: deviceRepo,
	}
}

func (u *DeviceUseCase) FindByID(
	ctx context.Context,
	req *apprequest.FindDeviceByIDRequest,
) (*appresponse.DeviceResponse, error) {

	device, err := u.deviceRepo.FindByID(
		ctx,
		req.DeviceID,
	)

	if err != nil {
		return nil, err
	}

	if device == nil {
		return nil, errordomain.ErrDeviceNotFound
	}

	return &appresponse.DeviceResponse{
		ID:   device.ID,
		Name: device.Name,
	}, nil
}
