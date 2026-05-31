package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/application/usecase"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation"
)

type DeviceController struct {
	deviceUseCase usecase.DeviceUseCaseInterface
}

func NewDeviceController(
	deviceUseCase usecase.DeviceUseCaseInterface,
) *DeviceController {
	return &DeviceController{
		deviceUseCase: deviceUseCase,
	}
}

func (c *DeviceController) FindByID(ctx *gin.Context) {

	req := new(request.FindDeviceByIDRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	if err := validation.Validate(
		validation.FindDeviceByIDRules(req),
	); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.deviceUseCase.FindByID(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		ctx.Error(response.MapError(err))
		return
	}

	response.ResponseOK(
		ctx,
		res,
	)
}
