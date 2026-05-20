package usecase

import (
	"context"

	apprequest "github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	appresponse "github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"

	"github.com/arthurhzna/go-clean-architecture/internal/application/validation/builder"
	"github.com/arthurhzna/go-clean-architecture/internal/application/validation/rule"

	errordomain "github.com/arthurhzna/go-clean-architecture/internal/domain/error"

	repositoryiface "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"
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

	err := rule.Execute(
		builder.FindDeviceByIDRules(*req),
	)

	if err != nil {
		return nil, err
	}

	device, err := u.deviceRepo.FindByID(
		ctx,
		req.ID,
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
