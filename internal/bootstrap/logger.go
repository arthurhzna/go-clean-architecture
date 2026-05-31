package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/config"
	loggerdomain "github.com/arthurhzna/go-clean-architecture/internal/domain/logger"
	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/logging"
)

func NewLogger(
	cfg *config.Config,
) loggerdomain.Logger {

	return logging.NewZeroLogger(
		cfg.Logger.Level,
	)
}
