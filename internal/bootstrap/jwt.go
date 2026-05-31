package bootstrap

import (
	"time"

	"github.com/arthurhzna/go-clean-architecture/internal/config"
	jwtinfra "github.com/arthurhzna/go-clean-architecture/internal/infrastructure/security/jwt"
)

func NewTokenService(
	cfg *config.Config,
) *jwtinfra.TokenService {

	return jwtinfra.NewTokenService(
		cfg.Jwt.SecretKey,
		cfg.Jwt.Issuer,
		time.Duration(cfg.Jwt.TokenDuration)*time.Minute,
		cfg.Jwt.AllowedAlgs,
	)
}
