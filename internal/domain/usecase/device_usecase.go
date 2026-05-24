package usecase

import (
	"context"

	apprequest "github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	appresponse "github.com/arthurhzna/go-clean-architecture/internal/application/dto/response"
)

type DeviceUseCase interface {
	FindByID(
		ctx context.Context,
		req *apprequest.FindDeviceByIDRequest,
	) (*appresponse.DeviceResponse, error)
}
