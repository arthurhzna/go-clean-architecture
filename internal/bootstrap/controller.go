package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/controller"
)

type Controller struct {
	AppController    *controller.AppController
	UserController   *controller.UserController
	DeviceController *controller.DeviceController
}

func NewController(
	useCase *UseCase,
) *Controller {
	return &Controller{
		AppController: controller.NewAppController(),

		UserController: controller.NewUserController(
			useCase.UserUseCase,
		),

		DeviceController: controller.NewDeviceController(
			useCase.DeviceUseCase,
		),
	}
}
