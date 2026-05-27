package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/config"

	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/logging"

	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/persistence/database"

	bcryptinfra "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/security/bcrypt"

	serviceinfra "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/service"
)

type Application struct {
	HttpServer *HttpServer
}

func NewApplication() *Application {

	// config
	cfg := config.InitConfig()

	// logger
	log := logging.NewZeroLogger(cfg)

	// database
	db := NewDatabase(
		cfg,
		log,
	)

	// unit of work
	uow := database.NewUnitOfWork(
		db,
	)

	// infrastructure
	jwtUtil := NewJwtUtil(cfg)

	passwordHasher := bcryptinfra.NewBcryptEncryptor(cfg.App.BCryptCost)

	uuidGenerator := serviceinfra.NewUUIDGenerator()

	// usecase
	useCase := NewUseCase(
		uow,
		passwordHasher,
		jwtUtil,
		uuidGenerator,
	)

	// controller
	controller := NewController(
		useCase,
	)

	// http server
	httpServer := NewHTTPServer(
		cfg,
		log,
		jwtUtil,
		controller,
	)

	return &Application{
		HttpServer: httpServer,
	}
}
