package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/config"

	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/persistence/database"

	bcryptinfra "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/security/bcrypt"

	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/identity"
)

type Application struct {
	HttpServer *HttpServer
}

func NewApplication() *Application {

	cfg := config.InitConfig()

	log := NewLogger(cfg)

	db := NewDatabase(
		cfg,
		log,
	)

	uow := database.NewUnitOfWork(
		db,
	)

	jwtUtil := NewTokenService(cfg)

	passwordHasher := bcryptinfra.NewBcryptEncryptor(cfg.App.BCryptCost)

	uuidGenerator := identity.NewUUIDGenerator()

	useCase := NewUseCase(
		uow,
		passwordHasher,
		jwtUtil,
		uuidGenerator,
	)

	controller := NewController(
		useCase,
	)

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
