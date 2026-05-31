package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/application/usecase"

	securitydomain "github.com/arthurhzna/go-clean-architecture/internal/domain/security"
	servicedomain "github.com/arthurhzna/go-clean-architecture/internal/domain/service"

	repositorydomain "github.com/arthurhzna/go-clean-architecture/internal/domain/repository"
)

type UseCase struct {
	UserUseCase   usecase.UserUseCaseInterface
	DeviceUseCase usecase.DeviceUseCaseIterface
}

func NewUseCase(
	uow repositorydomain.UnitOfWork,

	passwordHasher securitydomain.PasswordHasher,
	tokenService securitydomain.TokenService,
	uuidGenerator servicedomain.UUIDGenerator,
) *UseCase {

	return &UseCase{
		UserUseCase: usecase.NewUserUseCase(
			uow,
			passwordHasher,
			tokenService,
			uuidGenerator,
		),

		DeviceUseCase: usecase.NewDeviceUseCase(
			uow.DeviceRepository(),
		),
	}
}
