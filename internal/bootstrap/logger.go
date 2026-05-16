package bootstrap

import (
	"github.com/arthurhzna/go-clean-architecture/internal/config"
	"github.com/arthurhzna/go-clean-architecture/internal/infrastructure/logging"
)

func NewLogger(
	cfg *config.Config,
) *logging.ZeroLogger {

	return logging.NewZeroLogger(
		cfg.Logger.Level,
	)
}
